package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type sseHub struct {
	mu   sync.Mutex
	subs map[chan []byte]struct{}
}

func newHub() *sseHub { return &sseHub{subs: map[chan []byte]struct{}{}} }

func (h *sseHub) subscribe() chan []byte {
	ch := make(chan []byte, 64)
	h.mu.Lock()
	h.subs[ch] = struct{}{}
	h.mu.Unlock()
	return ch
}

func (h *sseHub) unsubscribe(ch chan []byte) {
	h.mu.Lock()
	delete(h.subs, ch)
	h.mu.Unlock()
	close(ch)
}

func (h *sseHub) send(v any) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	for ch := range h.subs {
		select {
		case ch <- b:
		default:
		}
	}
}

type guiState struct {
	mu      sync.Mutex
	running bool
	cancel  context.CancelFunc
	images  map[int]CollectedImage
	outDir  string
}

func runGUI(web fs.FS) error {
	hub := newHub()
	st := &guiState{images: map[int]CollectedImage{}}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(web)))
	mux.HandleFunc("/api/defaults", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{"outDir": "archived-images"})
	})
	mux.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {
		fl, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "no flush", 500)
			return
		}
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		ch := hub.subscribe()
		defer hub.unsubscribe(ch)
		notify := r.Context().Done()
		for {
			select {
			case <-notify:
				return
			case b, ok := <-ch:
				if !ok {
					return
				}
				fmt.Fprintf(w, "data: %s\n\n", b)
				fl.Flush()
			}
		}
	})
	mux.HandleFunc("/api/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", 405)
			return
		}
		var req struct {
			URL       string `json:"url"`
			OutDir    string `json:"outDir"`
			MaxImages int    `json:"maxImages"`
			ScanHTML  bool   `json:"scanHtml"`
			HTMLPages int    `json:"htmlPages"`
			DelayMs   int    `json:"delayMs"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad json", 400)
			return
		}
		st.mu.Lock()
		if st.running {
			st.mu.Unlock()
			http.Error(w, "already running", 409)
			return
		}
		st.running = true
		st.images = map[int]CollectedImage{}
		st.outDir = req.OutDir
		ctx, cancel := context.WithCancel(context.Background())
		st.cancel = cancel
		st.mu.Unlock()

		cfg := defaultConfig()
		cfg.StartURL = req.URL
		cfg.OutDir = req.OutDir
		if req.MaxImages > 0 {
			cfg.MaxImages = req.MaxImages
		}
		cfg.ScanHTML = req.ScanHTML
		if req.HTMLPages > 0 {
			cfg.HTMLPages = req.HTMLPages
		}
		if req.DelayMs >= 0 {
			cfg.DelayMs = req.DelayMs
		}
		cfg.Sink = func(img CollectedImage) {
			st.mu.Lock()
			st.images[img.ID] = img
			st.mu.Unlock()
			hub.send(map[string]any{
				"type": "image", "id": img.ID, "url": img.URL, "page": img.Page,
				"via": img.Via, "hidden": img.Hidden, "bytes": img.Bytes,
				"name": img.Name, "preview": img.Preview, "contentType": img.ContentType,
			})
		}
		go func() {
			defer func() {
				st.mu.Lock()
				st.running = false
				st.cancel = nil
				st.mu.Unlock()
			}()
			log := func(kind, msg string) {
				hub.send(map[string]any{"type": "log", "kind": kind, "msg": msg, "at": time.Now().Format("15:04:05")})
			}
			stats, err := Run(ctx, cfg, log)
			if err != nil {
				hub.send(map[string]any{"type": "error", "msg": err.Error()})
				hub.send(map[string]any{"type": "done", "ok": false})
				return
			}
			hub.send(map[string]any{
				"type": "done", "ok": true,
				"pages": stats.Pages.Load(), "found": stats.Found.Load(),
				"saved": stats.Saved.Load(), "hidden": stats.Hidden.Load(),
				"errors": stats.Errors.Load(),
			})
		}()
		writeJSON(w, map[string]any{"ok": true})
	})
	mux.HandleFunc("/api/stop", func(w http.ResponseWriter, r *http.Request) {
		st.mu.Lock()
		if st.cancel != nil {
			st.cancel()
		}
		st.mu.Unlock()
		writeJSON(w, map[string]any{"ok": true})
	})
	mux.HandleFunc("/api/image/", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/image/"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
		st.mu.Lock()
		img, ok := st.images[id]
		st.mu.Unlock()
		if !ok || len(img.Data) == 0 {
			http.NotFound(w, r)
			return
		}
		ct := img.ContentType
		if ct == "" {
			ct = "application/octet-stream"
		}
		w.Header().Set("Content-Type", ct)
		w.Write(img.Data)
	})
	mux.HandleFunc("/api/save", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", 405)
			return
		}
		var req struct {
			OutDir string `json:"outDir"`
			IDs    []int  `json:"ids"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad json", 400)
			return
		}
		st.mu.Lock()
		out := req.OutDir
		if out == "" {
			out = st.outDir
		}
		if out == "" {
			out = "archived-images"
		}
		var picks []CollectedImage
		if len(req.IDs) == 0 {
			for _, img := range st.images {
				picks = append(picks, img)
			}
		} else {
			for _, id := range req.IDs {
				if img, ok := st.images[id]; ok {
					picks = append(picks, img)
				}
			}
		}
		st.mu.Unlock()
		if err := os.MkdirAll(out, 0o755); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		n := 0
		var man []ManifestEntry
		for _, img := range picks {
			hit := ImageHit{URL: img.URL, Page: img.Page, Via: img.Via, Hidden: img.Hidden}
			name, err := writeImageFile(out, hit, img.Data, img.ContentType)
			if err != nil {
				continue
			}
			n++
			man = append(man, ManifestEntry{File: name, URL: img.URL, Page: img.Page, Via: img.Via, Hidden: img.Hidden, Bytes: img.Bytes})
		}
		if data, err := json.MarshalIndent(man, "", "  "); err == nil {
			_ = os.WriteFile(filepath.Join(out, "_manifest.json"), data, 0o644)
		}
		writeJSON(w, map[string]any{"ok": true, "saved": n, "outDir": out})
	})
	mux.HandleFunc("/api/quit", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{"ok": true})
		go func() {
			time.Sleep(200 * time.Millisecond)
			os.Exit(0)
		}()
	})

	ln, err := net.Listen("tcp", "127.0.0.1:47822")
	if err != nil {
		ln, err = net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			return err
		}
	}
	u := "http://" + ln.Addr().String() + "/"
	fmt.Println("Lumina Archive")
	fmt.Println("Search the Wayback Machine for every image a site published.")
	fmt.Println("Open this page if the browser does not appear:")
	fmt.Println("  " + u)
	openBrowser(u)
	return http.Serve(ln, mux)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func openBrowser(u string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", u)
	case "darwin":
		cmd = exec.Command("open", u)
	default:
		cmd = exec.Command("xdg-open", u)
	}
	_ = cmd.Start()
}
