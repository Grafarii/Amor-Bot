package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestHostileSiteCollectsEveryPublishedImage(t *testing.T) {
	var cdnHits, originHits int
	var mu sync.Mutex

	cdn := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		cdnHits++
		mu.Unlock()
		if r.Header.Get("Sec-Fetch-Site") == "same-origin" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}))
	defer cdn.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		originHits++
		mu.Unlock()
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `<!doctype html><html><head>
<link rel="stylesheet" href="/gone.css">
<meta property="og:image" content="/og.png">
</head><body>
<img src="/plain.png">
<img src="/hotlink.png">
<img src="/gate.png">
<img src="%s/cdn.png">
<img src="/extless">
<div style="background-image:url(/bg.png)"></div>
<a href="/album/">album</a>
<a href="/walled">walled</a>
</body></html>`, cdn.URL)
	})
	mux.HandleFunc("/gone.css", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	mux.HandleFunc("/walled", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	mux.HandleFunc("/album/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><img src="/album/shot.png"></body></html>`))
	})
	mux.HandleFunc("/hotlink.png", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Referer") == "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	})
	mux.HandleFunc("/gate.png", func(w http.ResponseWriter, r *http.Request) {
		if c, err := r.Cookie("gate"); err != nil || c.Value != "1" {
			http.SetCookie(w, &http.Cookie{Name: "gate", Value: "1", Path: "/"})
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	})
	mux.HandleFunc("/extless", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	})
	servePNG := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}
	for _, p := range []string{"/plain.png", "/og.png", "/bg.png", "/album/shot.png"} {
		mux.HandleFunc(p, servePNG)
	}

	origin := httptest.NewServer(mux)
	defer origin.Close()

	var got []CollectedImage
	var errLogs []string
	var logMu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = origin.URL + "/"
	cfg.SameHost = true
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Concurrency = 4
	cfg.MaxDepth = 2
	cfg.MaxPages = 20
	cfg.Sink = func(img CollectedImage) {
		logMu.Lock()
		got = append(got, img)
		logMu.Unlock()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	stats, err := Run(ctx, cfg, func(kind, msg string) {
		if kind == "error" {
			logMu.Lock()
			errLogs = append(errLogs, msg)
			logMu.Unlock()
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	if stats.Errors.Load() != 0 {
		t.Fatalf("hostile site must not stop with errors=%d logs=%v", stats.Errors.Load(), errLogs)
	}
	if len(errLogs) != 0 {
		t.Fatalf("error log not empty: %v", errLogs)
	}

	joined := ""
	logMu.Lock()
	for _, img := range got {
		joined += img.URL + " "
	}
	logMu.Unlock()
	need := []string{"plain.png", "hotlink.png", "gate.png", "cdn.png", "og.png", "bg.png", "shot.png", "extless"}
	var missing []string
	for _, n := range need {
		if !strings.Contains(joined, n) {
			missing = append(missing, n)
		}
	}
	if len(missing) > 0 {
		t.Fatalf("did not collect %v from %s (cdnHits=%d originHits=%d pages=%d)", missing, joined, cdnHits, originHits, stats.Pages.Load())
	}
}
