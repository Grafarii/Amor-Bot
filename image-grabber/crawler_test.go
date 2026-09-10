package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// 1x1 transparent PNG
var pngDot = []byte{
	0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
	0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
	0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
	0x08, 0x06, 0x00, 0x00, 0x00, 0x1F, 0x15, 0xC4, 0x89,
	0x00, 0x00, 0x00, 0x0A, 0x49, 0x44, 0x41, 0x54,
	0x78, 0x9C, 0x63, 0x00, 0x01, 0x00, 0x00, 0x05, 0x00, 0x01,
	0x0D, 0x0A, 0x2D, 0xB4, 0x00, 0x00, 0x00, 0x00,
	0x49, 0x45, 0x4E, 0x44, 0xAE, 0x42, 0x60, 0x82,
}

func TestCrawlHiddenIndexAndAssets(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("User-agent: *\nAllow: /\nSitemap: http://" + r.Host + "/sitemap.xml\n"))
	})
	mux.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(`<?xml version="1.0"?><urlset><url><loc>http://` + r.Host + `/</loc><image:loc>http://` + r.Host + `/from-sitemap.png</image:loc></url></urlset>`))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!doctype html>
<html><head>
<link rel="stylesheet" href="/app.css">
<script src="/app.js"></script>
</head><body>
<img src="/visible.png">
<!-- <img src="/commented.png"> -->
<div hidden><img src="/hidden.png"></div>
<a href="/index/">photos</a>
</body></html>`))
	})
	mux.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(`body{background:url(/css.png)}`))
	})
	mux.HandleFunc("/app.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(`var x = "/js.png";`))
	})
	mux.HandleFunc("/index/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><h1>Index of /index/</h1><a href="/listed.png">listed.png</a></body></html>`))
	})
	servePNG := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}
	for _, p := range []string{"/visible.png", "/commented.png", "/hidden.png", "/css.png", "/js.png", "/listed.png", "/from-sitemap.png"} {
		mux.HandleFunc(p, servePNG)
	}

	srv := httptest.NewServer(mux)
	defer srv.Close()

	out := t.TempDir()
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.OutDir = out
	cfg.MaxDepth = 2
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.Concurrency = 4

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	stats, err := Run(ctx, cfg, func(string, string) {})
	if err != nil {
		t.Fatal(err)
	}
	if stats.Saved.Load() < 7 {
		t.Fatalf("expected at least 7 saved images, got %d (found=%d errors=%d)", stats.Saved.Load(), stats.Found.Load(), stats.Errors.Load())
	}
	if stats.Hidden.Load() < 2 {
		t.Fatalf("expected hidden images, got %d", stats.Hidden.Load())
	}

	raw, err := os.ReadFile(filepath.Join(out, "_manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var man []ManifestEntry
	if err := json.Unmarshal(raw, &man); err != nil {
		t.Fatal(err)
	}
	joined := ""
	for _, e := range man {
		joined += e.URL + " " + e.Via + "\n"
	}
	for _, needle := range []string{"visible.png", "commented.png", "hidden.png", "css.png", "js.png", "listed.png", "from-sitemap.png"} {
		if !strings.Contains(joined, needle) {
			t.Errorf("manifest missing %s\n%s", needle, joined)
		}
	}
}

func TestPreviewSinkDoesNotWriteFiles(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><img src="/visible.png"><div hidden><img src="/hidden.png"></div></body></html>`))
	})
	servePNG := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}
	mux.HandleFunc("/visible.png", servePNG)
	mux.HandleFunc("/hidden.png", servePNG)

	srv := httptest.NewServer(mux)
	defer srv.Close()

	out := t.TempDir()
	var got []CollectedImage
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.OutDir = out
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Sink = func(img CollectedImage) {
		mu.Lock()
		got = append(got, img)
		mu.Unlock()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stats, err := Run(ctx, cfg, func(string, string) {})
	if err != nil {
		t.Fatal(err)
	}
	if stats.Saved.Load() != 0 {
		t.Fatalf("preview mode should not save, got %d", stats.Saved.Load())
	}
	entries, _ := os.ReadDir(out)
	if len(entries) != 0 {
		t.Fatalf("preview mode wrote files: %v", entries)
	}
	mu.Lock()
	n := len(got)
	mu.Unlock()
	if n < 2 {
		t.Fatalf("expected collected images, got %d", n)
	}
}

func TestForbiddenIsSkippedNotError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><img src="/ok.png"><img src="/denied.png"></body></html>`))
	})
	mux.HandleFunc("/ok.png", func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("User-Agent"), "Chrome") {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	})
	mux.HandleFunc("/denied.png", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><h1>Index of /users/</h1><img src="/users/avatar.png" data-avatar="/users/avatar.png"></body></html>`))
	})
	mux.HandleFunc("/users/avatar.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	var got []CollectedImage
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Sink = func(img CollectedImage) {
		mu.Lock()
		got = append(got, img)
		mu.Unlock()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	stats, err := Run(ctx, cfg, func(string, string) {})
	if err != nil {
		t.Fatal(err)
	}
	if stats.Errors.Load() != 0 {
		t.Fatalf("HTTP 403 should be skipped, not an error; errors=%d skipped=%d", stats.Errors.Load(), stats.Skipped.Load())
	}
	joined := ""
	mu.Lock()
	for _, img := range got {
		joined += img.URL + " "
	}
	mu.Unlock()
	if !strings.Contains(joined, "ok.png") {
		t.Fatalf("missing ok.png (browser UA / referer). got %s", joined)
	}
	if !strings.Contains(joined, "avatar.png") {
		t.Fatalf("missing users avatar. got %s", joined)
	}
}

func TestCrawlCollectsMp4AndOddImageExt(t *testing.T) {
	mp4 := []byte{0x00, 0x00, 0x00, 0x18, 'f', 't', 'y', 'p', 'i', 's', 'o', 'm', 0x00, 0x00, 0x02, 0x00, 'i', 's', 'o', 'm'}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body>
<video src="/clip.mp4" poster="/poster.jpg"><source src="/alt.webm"></video>
<a href="/scan.tiff">tiff</a>
<img src="/pic.bmp">
</body></html>`))
	})
	mux.HandleFunc("/clip.mp4", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "video/mp4")
		w.Write(mp4)
	})
	mux.HandleFunc("/alt.webm", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(mp4)
	})
	mux.HandleFunc("/poster.jpg", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write([]byte{0xff, 0xd8, 0xff, 0xd9})
	})
	mux.HandleFunc("/scan.tiff", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/tiff")
		w.Write([]byte("II*\x00not-a-real-tiff-but-typed"))
	})
	mux.HandleFunc("/pic.bmp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/bmp")
		w.Write([]byte("BM" + strings.Repeat("x", 20)))
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	var got []CollectedImage
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Sink = func(img CollectedImage) {
		mu.Lock()
		got = append(got, img)
		mu.Unlock()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if _, err := Run(ctx, cfg, func(string, string) {}); err != nil {
		t.Fatal(err)
	}
	joined := ""
	mu.Lock()
	for _, img := range got {
		joined += img.URL + " "
	}
	mu.Unlock()
	for _, needle := range []string{"clip.mp4", "alt.webm", "poster.jpg", "scan.tiff", "pic.bmp"} {
		if !strings.Contains(joined, needle) {
			t.Errorf("missing %s in %s", needle, joined)
		}
	}
}
