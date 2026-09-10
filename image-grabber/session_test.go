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

func TestCrawlReads403HTMLAndKeepsWalking(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`<html><body><img src="/hero.png"><a href="/album/">album</a></body></html>`))
	})
	mux.HandleFunc("/album/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><img src="/shot.png"></body></html>`))
	})
	servePNG := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	}
	mux.HandleFunc("/hero.png", servePNG)
	mux.HandleFunc("/shot.png", servePNG)
	srv := httptest.NewServer(mux)
	defer srv.Close()

	var got []CollectedImage
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Concurrency = 1
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
		t.Fatalf("403 html must not stop the walk with errors=%d", stats.Errors.Load())
	}
	joined := ""
	mu.Lock()
	for _, img := range got {
		joined += img.URL + " "
	}
	mu.Unlock()
	if !strings.Contains(joined, "hero.png") || !strings.Contains(joined, "shot.png") {
		t.Fatalf("expected walk to continue past 403 html, got %s", joined)
	}
}

func Test403ImageBodyIsCollected(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<html><body><img src="/hot.png"></body></html>`))
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusForbidden)
		w.Write(pngDot)
	}))
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := Run(ctx, cfg, func(string, string) {}); err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	defer mu.Unlock()
	if len(got) < 1 {
		t.Fatal("expected 403 image body to be collected")
	}
}

func TestForbiddenPagesDoNotBurnPageBudget(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><a href="/wall">wall</a><a href="/good">good</a></body></html>`))
	})
	mux.HandleFunc("/wall", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
	mux.HandleFunc("/good", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<html><body><img src="/kept.png"></body></html>`))
	})
	mux.HandleFunc("/kept.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(pngDot)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	var got []CollectedImage
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.MaxPages = 2
	cfg.MaxDepth = 2
	cfg.DelayMs = 0
	cfg.Concurrency = 1
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Sink = func(img CollectedImage) {
		mu.Lock()
		got = append(got, img)
		mu.Unlock()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	if !strings.Contains(joined, "kept.png") {
		t.Fatalf("403 page burned the page budget; missing kept.png in %s", joined)
	}
}
