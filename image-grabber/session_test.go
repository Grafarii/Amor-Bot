package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSecFetchSite(t *testing.T) {
	must := func(raw string) *url.URL {
		u, err := url.Parse(raw)
		if err != nil {
			t.Fatal(err)
		}
		return u
	}
	if got := secFetchSite(must("https://example.com/a"), must("https://example.com/b")); got != "same-origin" {
		t.Fatalf("same origin: %s", got)
	}
	if got := secFetchSite(must("https://example.com/"), must("https://cdn.example.com/x.png")); got != "same-site" {
		t.Fatalf("cdn subdomain: %s", got)
	}
	if got := secFetchSite(must("https://site.test/"), must("https://images.elsewhere.test/x.png")); got != "cross-site" {
		t.Fatalf("other host: %s", got)
	}
	if got := secFetchSite(must("http://127.0.0.1:8000/"), must("http://127.0.0.1:8001/x.png")); got != "same-site" {
		t.Fatalf("same host different port is same-site: %s", got)
	}
}

func TestGetURLForgivingCrossSiteHotlink(t *testing.T) {
	cdn := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Sec-Fetch-Site") != "cross-site" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if !strings.Contains(r.Header.Get("User-Agent"), "Chrome") {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}))
	defer cdn.Close()

	page := "http://gallery.test/"
	client := newCrawlClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body, ct, _, err := getURLForgiving(ctx, client, cdn.URL+"/hero.png", page, "media", page, 1<<20)
	if err != nil {
		t.Fatalf("expected CDN image, got %v", err)
	}
	if !strings.Contains(ct, "png") && len(body) < 8 {
		t.Fatalf("bad body ct=%s n=%d", ct, len(body))
	}
}

func TestGetURLForgivingNoRefererHotlink(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Referer") != "" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}))
	defer srv.Close()

	client := newCrawlClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _, _, err := getURLForgiving(ctx, client, srv.URL+"/pic.png", srv.URL+"/", "media", srv.URL+"/", 1<<20)
	if err != nil {
		t.Fatalf("expected image after dropping Referer, got %v", err)
	}
}

func TestGetURLForgivingCookieGate(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c, err := r.Cookie("gate"); err != nil || c.Value != "1" {
			http.SetCookie(w, &http.Cookie{Name: "gate", Value: "1", Path: "/"})
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}))
	defer srv.Close()

	client := newCrawlClient()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, _, _, err := getURLForgiving(ctx, client, srv.URL+"/pic.png", srv.URL+"/", "media", srv.URL+"/", 1<<20)
	if err != nil {
		t.Fatalf("expected image after cookie retry, got %v", err)
	}
}

func TestCrawlCollectsHotlinkProtectedCDN(t *testing.T) {
	var hits atomic.Int64
	cdn := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		site := r.Header.Get("Sec-Fetch-Site")
		ref := r.Header.Get("Referer")
		if site == "same-origin" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if ref == "" && site != "cross-site" && site != "same-site" && site != "none" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}))
	defer cdn.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><img src="` + cdn.URL + `/hero.png"><img src="/denied.png"></body></html>`))
	})
	mux.HandleFunc("/denied.png", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	origin := httptest.NewServer(mux)
	defer origin.Close()

	var got []CollectedImage
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = origin.URL + "/"
	cfg.SameHost = true
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
		t.Fatalf("403 must not be an error; errors=%d skipped=%d", stats.Errors.Load(), stats.Skipped.Load())
	}
	joined := ""
	mu.Lock()
	for _, img := range got {
		joined += img.URL + " "
	}
	mu.Unlock()
	if !strings.Contains(joined, "hero.png") {
		t.Fatalf("missing CDN hero after hotlink retries: %s (cdn hits=%d)", joined, hits.Load())
	}
	if strings.Contains(joined, "denied.png") {
		t.Fatalf("collected a permanently forbidden file: %s", joined)
	}
}
