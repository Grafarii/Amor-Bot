package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	StartURL   string
	OutDir     string
	MaxImages  int
	ScanHTML   bool
	HTMLPages  int
	DelayMs    int
	CDXBase    string
	Wayback    string
	Sink       func(CollectedImage) `json:"-"`
	HTTPClient *http.Client         `json:"-"`
}

type CollectedImage struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Page        string `json:"page"`
	Via         string `json:"via"`
	Hidden      bool   `json:"hidden"`
	Bytes       int    `json:"bytes"`
	ContentType string `json:"contentType"`
	Name        string `json:"name"`
	Preview     string `json:"preview"`
	Data        []byte `json:"-"`
}

type ManifestEntry struct {
	File   string `json:"file"`
	URL    string `json:"url,omitempty"`
	Page   string `json:"page,omitempty"`
	Via    string `json:"via,omitempty"`
	Hidden bool   `json:"hidden,omitempty"`
	Bytes  int    `json:"bytes,omitempty"`
}

type Stats struct {
	Pages  atomic.Int64
	Found  atomic.Int64
	Saved  atomic.Int64
	Hidden atomic.Int64
	Errors atomic.Int64
}

type LogFn func(kind, msg string)

func defaultConfig() Config {
	return Config{
		MaxImages: 800,
		ScanHTML:  true,
		HTMLPages: 40,
		DelayMs:   160,
		CDXBase:   defaultCDX,
		Wayback:   defaultWayb,
	}
}

func newArchiveClient() *http.Client {
	return &http.Client{Timeout: 75 * time.Second}
}

func Run(ctx context.Context, cfg Config, log LogFn) (*Stats, error) {
	if log == nil {
		log = func(string, string) {}
	}
	start, err := url.Parse(strings.TrimSpace(cfg.StartURL))
	if err != nil || start.Host == "" {
		return nil, fmt.Errorf("invalid URL: need a full address like https://example.com")
	}
	if start.Scheme == "" {
		start.Scheme = "https"
	}
	preview := cfg.Sink != nil
	if !preview {
		if cfg.OutDir == "" {
			cfg.OutDir = "archived-images"
		}
		if err := os.MkdirAll(cfg.OutDir, 0o755); err != nil {
			return nil, err
		}
	}
	if cfg.MaxImages < 1 {
		cfg.MaxImages = 200
	}
	if cfg.HTMLPages < 1 {
		cfg.HTMLPages = 20
	}
	if cfg.CDXBase == "" {
		cfg.CDXBase = defaultCDX
	}
	if cfg.Wayback == "" {
		cfg.Wayback = defaultWayb
	}
	client := cfg.HTTPClient
	if client == nil {
		client = newArchiveClient()
	}
	delay := time.Duration(cfg.DelayMs) * time.Millisecond
	stats := &Stats{}
	target, match := siteQuery(start)
	log("info", "searching the Wayback Machine for "+target)

	pace := func() {
		if delay > 0 {
			select {
			case <-ctx.Done():
			case <-time.After(delay):
			}
		}
	}

	var (
		mu       sync.Mutex
		seenURL  = map[string]bool{}
		seenDig  = map[string]bool{}
		manifest []ManifestEntry
		nextID   atomic.Int64
		left     = int64(cfg.MaxImages)
	)

	save := func(hit ImageHit, body []byte, ct string) {
		if len(body) == 0 || !isMediaContent(ct, body) && !looksLikeImage(hit.URL) {
			if len(body) == 0 || !isMediaContent(ct, body) {
				return
			}
		}
		if atomic.AddInt64(&left, -1) < 0 {
			atomic.AddInt64(&left, 1)
			return
		}
		stats.Found.Add(1)
		name := suggestedName(hit, body, ct)
		if preview {
			if hit.Hidden {
				stats.Hidden.Add(1)
			}
			id := int(nextID.Add(1))
			cfg.Sink(CollectedImage{
				ID: id, URL: hit.URL, Page: hit.Page, Via: hit.Via, Hidden: hit.Hidden,
				Bytes: len(body), ContentType: ct, Name: filepath.ToSlash(name),
				Preview: fmt.Sprintf("/api/image/%d", id), Data: body,
			})
			log("found", filepath.Base(name)+" ← "+hit.Via)
			return
		}
		written, err := writeImageFile(cfg.OutDir, hit, body, ct)
		if err != nil {
			stats.Errors.Add(1)
			return
		}
		stats.Saved.Add(1)
		if hit.Hidden {
			stats.Hidden.Add(1)
		}
		mu.Lock()
		manifest = append(manifest, ManifestEntry{File: written, URL: hit.URL, Page: hit.Page, Via: hit.Via, Hidden: hit.Hidden, Bytes: len(body)})
		mu.Unlock()
		log("save", filepath.Base(written)+" ← "+hit.Via)
	}

	fetchSnap := func(rec CDXRecord, via, page string) {
		if ctx.Err() != nil {
			return
		}
		mu.Lock()
		if rec.Digest != "" && seenDig[rec.Digest] {
			mu.Unlock()
			return
		}
		key := strings.ToLower(rec.Original)
		if seenURL[key] {
			mu.Unlock()
			return
		}
		seenURL[key] = true
		if rec.Digest != "" {
			seenDig[rec.Digest] = true
		}
		mu.Unlock()

		try := []string{
			snapshotURL(cfg.Wayback, rec.Timestamp, rec.Original),
			latestSnapshotURL(cfg.Wayback, rec.Original),
		}
		var body []byte
		var ct string
		ok := false
		for _, snap := range try {
			if snap == "" || ctx.Err() != nil {
				continue
			}
			pace()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, snap, nil)
			if err != nil {
				continue
			}
			req.Header.Set("User-Agent", archiveUA)
			req.Header.Set("Accept", "*/*")
			resp, err := client.Do(req)
			if err != nil {
				continue
			}
			raw, err := io.ReadAll(io.LimitReader(resp.Body, 20<<20))
			resp.Body.Close()
			if err != nil || resp.StatusCode >= 400 || len(raw) == 0 {
				continue
			}
			body, ct, ok = raw, resp.Header.Get("Content-Type"), true
			break
		}
		if !ok {
			log("skip", rec.Original)
			return
		}
		save(ImageHit{URL: rec.Original, Page: page, Via: via}, body, ct)
	}

	collectIndex := func(label string, filters []string, limit int) {
		if ctx.Err() != nil {
			return
		}
		pace()
		recs, err := queryCDX(ctx, client, cfg.CDXBase, target, match, filters, limit)
		if err != nil {
			log("skip", label+": "+err.Error())
			return
		}
		log("scan", fmt.Sprintf("%d %s in the CDX index", len(recs), label))
		for _, rec := range recs {
			if ctx.Err() != nil {
				return
			}
			if atomic.LoadInt64(&left) <= 0 {
				return
			}
			fetchSnap(rec, "wayback-cdx", rec.Original)
		}
	}

	collectIndex("archived images", []string{"mimetype:image/.*"}, cfg.MaxImages)
	if atomic.LoadInt64(&left) > 0 {
		collectIndex("archived video", []string{"mimetype:video/.*"}, min(400, cfg.MaxImages))
	}
	if atomic.LoadInt64(&left) > 0 {
		collectIndex("files with media names", []string{extFilter}, cfg.MaxImages)
	}

	if cfg.ScanHTML && ctx.Err() == nil {
		pace()
		pages, err := queryCDX(ctx, client, cfg.CDXBase, target, match, []string{"mimetype:text/html"}, cfg.HTMLPages)
		if err != nil {
			log("skip", "page index: "+err.Error())
		} else {
			log("scan", fmt.Sprintf("reading %d archived pages for buried images", len(pages)))
			for _, rec := range pages {
				if ctx.Err() != nil {
					break
				}
				if atomic.LoadInt64(&left) <= 0 {
					break
				}
				snap := snapshotURL(cfg.Wayback, rec.Timestamp, rec.Original)
				pace()
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, snap, nil)
				if err != nil {
					continue
				}
				req.Header.Set("User-Agent", archiveUA)
				req.Header.Set("Accept", "text/html")
				resp, err := client.Do(req)
				if err != nil {
					continue
				}
				body, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
				resp.Body.Close()
				if err != nil || resp.StatusCode >= 400 {
					continue
				}
				stats.Pages.Add(1)
				pageURL, _ := url.Parse(rec.Original)
				if pageURL == nil {
					continue
				}
				log("page", rec.Original)
				res := extractHTML(pageURL, body)
				for _, img := range res.Images {
					if ctx.Err() != nil {
						break
					}
					if len(img.Data) > 0 {
						ct := extToMIME(img.Ext)
						if ct == "" {
							ct = "application/octet-stream"
						}
						save(img, img.Data, ct)
						continue
					}
					orig := unwrapArchiveURL(img.URL)
					if orig == "" || !looksLikeImage(orig) {
						continue
					}
					mu.Lock()
					if seenURL[strings.ToLower(orig)] {
						mu.Unlock()
						continue
					}
					mu.Unlock()
					fetchSnap(CDXRecord{Timestamp: rec.Timestamp, Original: orig}, "wayback-html", rec.Original)
				}
			}
		}
	}

	if !preview && len(manifest) > 0 {
		if data, err := json.MarshalIndent(manifest, "", "  "); err == nil {
			_ = os.WriteFile(filepath.Join(cfg.OutDir, "_manifest.json"), data, 0o644)
		}
	}
	log("info", fmt.Sprintf("done — %d archive pages, %d images", stats.Pages.Load(), stats.Found.Load()))
	return stats, nil
}

func unwrapArchiveURL(raw string) string {
	orig := strings.TrimSpace(raw)
	if i := strings.Index(orig, "id_/"); i >= 0 {
		return orig[i+4:]
	}
	if i := strings.Index(orig, "im_/"); i >= 0 {
		return orig[i+4:]
	}
	if i := strings.Index(orig, "if_/"); i >= 0 {
		return orig[i+4:]
	}
	if j := strings.Index(orig, "/web/"); j >= 0 {
		rest := orig[j+5:]
		if k := strings.Index(rest, "/http"); k >= 0 {
			return rest[k+1:]
		}
		if k := strings.Index(rest, "/https"); k >= 0 {
			return rest[k+1:]
		}
	}
	return orig
}
