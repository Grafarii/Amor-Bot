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
	"sync/atomic"
	"time"
)

type sseHub struct {
	mu   sync.Mutex
	subs map[chan []byte]struct{}
}

func newHub() *sseHub {
	return &sseHub{subs: map[chan []byte]struct{}{}}
}

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
	mu       sync.Mutex
	running  bool
	cancel   context.CancelFunc
	unlocked atomic.Bool
	images   map[int]CollectedImage
	outDir   string
}

type startReq struct {
	URL           string `json:"url"`
	OutDir        string `json:"outDir"`
	MaxDepth      int    `json:"maxDepth"`
	MaxPages      int    `json:"maxPages"`
	SameHost      bool   `json:"sameHost"`
	ParseSitemap  bool   `json:"parseSitemap"`
	ParseCSS      bool   `json:"parseCss"`
	ParseJS       bool   `json:"parseJs"`
	ParseComments bool   `json:"parseComments"`
	SaveDataURI   bool   `json:"saveDataUri"`
	RespectRobots bool   `json:"respectRobots"`
	DelayMs       int    `json:"delayMs"`
	Password      string `json:"password"`
	Cookies       string `json:"cookies"`
	LoginURL      string `json:"loginUrl"`
	LoginUser     string `json:"loginUser"`
	LoginPass     string `json:"loginPass"`
}

func runGUI(web fs.FS) error {
	hub := newHub()
	st := &guiState{images: map[int]CollectedImage{}}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(web)))
	mux.HandleFunc("/api/defaults", func(w http.ResponseWriter, r *http.Request) {
		exeDir, _ := os.Getwd()
		if p, err := os.Executable(); err == nil {
			exeDir = filepath.Dir(p)
		}
		writeJSON(w, map[string]any{
			"outDir":        filepath.Join(exeDir, "grabbed-images"),
			"unlocked":      st.unlocked.Load(),
			"safeMaxDepth":  safeMaxDepth,
			"safeMaxPages":  safeMaxPages,
			"safeMinDelay":  safeMinDelay,
			"adminMaxDepth": adminMaxDepth,
			"adminMaxPages": adminMaxPages,
		})
	})
	mux.HandleFunc("/api/events", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "no sse", 500)
			return
		}
		ch := hub.subscribe()
		defer hub.unsubscribe(ch)
		fmt.Fprintf(w, "data: {\"type\":\"hello\"}\n\n")
		flusher.Flush()
		for {
			select {
			case <-r.Context().Done():
				return
			case msg, ok := <-ch:
				if !ok {
					return
				}
				fmt.Fprintf(w, "data: %s\n\n", msg)
				flusher.Flush()
			}
		}
	})
	mux.HandleFunc("/api/unlock", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", 405)
			return
		}
		var req struct {
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad json", 400)
			return
		}
		if !passwordOK(req.Password) {
			http.Error(w, "wrong password", 401)
			return
		}
		st.unlocked.Store(true)
		writeJSON(w, map[string]any{"ok": true, "unlocked": true})
	})
	mux.HandleFunc("/api/start", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", 405)
			return
		}
		var req startReq
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad json", 400)
			return
		}
		unlocked := st.unlocked.Load() || passwordOK(req.Password)
		st.mu.Lock()
		if st.running {
			st.mu.Unlock()
			http.Error(w, "already running", 409)
			return
		}
		if unlocked {
			st.unlocked.Store(true)
		}
		st.images = map[int]CollectedImage{}
		st.outDir = req.OutDir
		ctx, cancel := context.WithCancel(context.Background())
		st.running = true
		st.cancel = cancel
		st.mu.Unlock()

		cfg := defaultConfig()
		cfg.StartURL = req.URL
		cfg.OutDir = req.OutDir
		if req.MaxDepth > 0 {
			cfg.MaxDepth = req.MaxDepth
		}
		if req.MaxPages > 0 {
			cfg.MaxPages = req.MaxPages
		}
		cfg.SameHost = req.SameHost
		cfg.ParseSitemap = req.ParseSitemap
		cfg.ParseCSS = req.ParseCSS
		cfg.ParseJS = req.ParseJS
		cfg.ParseComments = req.ParseComments
		cfg.SaveDataURI = req.SaveDataURI
		cfg.RespectRobots = req.RespectRobots
		if req.DelayMs >= 0 {
			cfg.DelayMs = req.DelayMs
		}
		cfg.Cookies = req.Cookies
		cfg.LoginURL = req.LoginURL
		cfg.LoginUser = req.LoginUser
		cfg.LoginPass = req.LoginPass
		boundMsg := applyBounds(&cfg, unlocked)
		cfg.Unlimited = &st.unlocked
		cfg.Sink = func(img CollectedImage) {
			st.mu.Lock()
			st.images[img.ID] = img
			st.mu.Unlock()
			hub.send(map[string]any{
				"type":        "image",
				"id":          img.ID,
				"url":         img.URL,
				"page":        img.Page,
				"via":         img.Via,
				"hidden":      img.Hidden,
				"bytes":       img.Bytes,
				"name":        img.Name,
				"preview":     img.Preview,
				"contentType": img.ContentType,
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
				if strings.Contains(msg, "queue full") && cfg.neverDropQueue() {
					return
				}
				hub.send(map[string]any{"type": "log", "kind": kind, "msg": msg, "at": time.Now().Format("15:04:05")})
			}
			log("info", boundMsg)
			stats, err := Run(ctx, cfg, log)
			if err != nil {
				hub.send(map[string]any{"type": "error", "msg": err.Error()})
				hub.send(map[string]any{"type": "done", "ok": false})
				return
			}
			hub.send(map[string]any{
				"type":   "done",
				"ok":     true,
				"pages":  stats.Pages.Load(),
				"found":  stats.Found.Load(),
				"saved":  stats.Saved.Load(),
				"hidden": stats.Hidden.Load(),
				"errors": stats.Errors.Load(),
				"outDir": cfg.OutDir,
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
		ct := previewContentType(img.ContentType, img.Name, img.URL)
		w.Header().Set("Content-Type", ct)
		w.Header().Set("Cache-Control", "private, max-age=60")
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
			out = "grabbed-images"
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
		var man []ManifestEntry
		n := 0
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

	ln, err := net.Listen("tcp", "127.0.0.1:47821")
	if err != nil {
		ln, err = net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			return err
		}
	}
	url := "http://" + ln.Addr().String() + "/"
	fmt.Println("Site Image Grabber")
	fmt.Println("Open this page if the browser does not appear:")
	fmt.Println("  " + url)
	fmt.Println("Testing limits are locked. Unlock in the UI with the admin password.")
	fmt.Println("Close this window or press Ctrl+C to quit.")
	openBrowser(url)
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
