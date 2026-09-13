package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"
)

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

func TestParseCDXJSON(t *testing.T) {
	raw := []byte(`[["timestamp","original","mimetype","statuscode","digest","length"],["20200101120000","http://example.com/a.png","image/png","200","AAA","12"]]`)
	recs, err := parseCDXJSON(raw)
	if err != nil || len(recs) != 1 {
		t.Fatalf("got %v %v", recs, err)
	}
	if recs[0].Original != "http://example.com/a.png" {
		t.Fatalf("original %q", recs[0].Original)
	}
}

func TestSnapshotURL(t *testing.T) {
	got := snapshotURL("https://web.archive.org", "20200101120000", "http://example.com/a.png")
	if got != "https://web.archive.org/web/20200101120000id_/http://example.com/a.png" {
		t.Fatalf("got %s", got)
	}
}

func TestSiteQuery(t *testing.T) {
	u, _ := url.Parse("https://www.Example.com/gallery")
	target, match := siteQuery(u)
	if target != "example.com/gallery/" || match != "prefix" {
		t.Fatalf("got %s %s", target, match)
	}
	root, _ := url.Parse("https://www.Example.com/")
	target, match = siteQuery(root)
	if target != "example.com" || match != "domain" {
		t.Fatalf("root got %s %s", target, match)
	}
}

func TestParseCDXText(t *testing.T) {
	raw := []byte("com,example)/a.png 20200101120000 http://example.com/a.png image/png 200 AAA 12")
	recs, err := parseCDXText(raw)
	if err != nil || len(recs) != 1 || recs[0].Original != "http://example.com/a.png" {
		t.Fatalf("got %v %v", recs, err)
	}
}

func TestUnwrapArchiveURL(t *testing.T) {
	got := unwrapArchiveURL("https://web.archive.org/web/20200101120000im_/http://example.com/a.png")
	if got != "http://example.com/a.png" {
		t.Fatalf("got %s", got)
	}
}

func TestCollectsCDXImagesAndHTMLBuried(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/cdx/search/cdx", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		filters := q["filter"]
		joined := strings.Join(filters, " ")
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(joined, "mimetype:image") {
			w.Write([]byte(`[["timestamp","original","mimetype","statuscode","digest","length"],["20200101120000","http://example.com/cdx.png","image/png","200","D1","20"]]`))
			return
		}
		w.Write([]byte(`[["timestamp","original","mimetype","statuscode","digest","length"],["20210101120000","http://example.com/old.html","text/html","200","H1","80"]]`))
	})
	mux.HandleFunc("/web/", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		if strings.Contains(p, "cdx.png") {
			w.Header().Set("Content-Type", "image/png")
			w.Write(pngDot)
			return
		}
		if strings.Contains(p, "old.html") {
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte(`<html><body><img src="http://example.com/buried.png"></body></html>`))
			return
		}
		if strings.Contains(p, "buried.png") {
			w.Header().Set("Content-Type", "image/png")
			w.Write(pngDot)
			return
		}
		http.NotFound(w, r)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	var got []CollectedImage
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = "https://example.com/"
	cfg.DelayMs = 0
	cfg.ScanHTML = true
	cfg.MaxImages = 20
	cfg.HTMLPages = 10
	cfg.CDXBase = srv.URL + "/cdx/search/cdx"
	cfg.Wayback = srv.URL
	cfg.HTTPClient = srv.Client()
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
	if stats.Errors.Load() != 0 {
		t.Fatalf("errors=%d", stats.Errors.Load())
	}
	joined := ""
	mu.Lock()
	for _, img := range got {
		joined += img.URL + " "
	}
	mu.Unlock()
	if !strings.Contains(joined, "cdx.png") || !strings.Contains(joined, "buried.png") {
		t.Fatalf("missing archive images in %s", joined)
	}
}
