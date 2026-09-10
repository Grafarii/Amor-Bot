package main

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const userAgent = browserUA

type Config struct {
	StartURL      string
	OutDir        string
	MaxDepth      int
	MaxPages      int
	SameHost      bool
	ParseSitemap  bool
	ParseCSS      bool
	ParseJS       bool
	ParseComments bool // kept for UI; comments are always parsed in HTML
	SaveDataURI   bool
	RespectRobots bool
	Concurrency   int
	DelayMs       int
	MaxBytes      int64
	Cookies       string
	LoginURL      string
	LoginUser     string
	LoginPass     string
	DeepScan      bool
	Unlimited     *atomic.Bool         `json:"-"`
	Sink          func(CollectedImage) `json:"-"`
}

func (c Config) neverDropQueue() bool {
	if c.Unlimited != nil && c.Unlimited.Load() {
		return true
	}
	return c.DeepScan
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

type LogFn func(kind, msg string)

type Stats struct {
	Pages   atomic.Int64
	Found   atomic.Int64
	Saved   atomic.Int64
	Skipped atomic.Int64
	Hidden  atomic.Int64
	Errors  atomic.Int64
}

type ManifestEntry struct {
	File   string `json:"file"`
	URL    string `json:"url,omitempty"`
	Page   string `json:"page"`
	Via    string `json:"via"`
	Hidden bool   `json:"hidden"`
	Bytes  int    `json:"bytes"`
}

type jobKind int

const (
	jobPage jobKind = iota
	jobStyle
	jobScript
	jobSitemap
	jobManifest
	jobImage
)

type job struct {
	kind   jobKind
	url    string
	depth  int
	via    string
	hidden bool
	page   string
	data   []byte
	ext    string
}

// hostBound is true when this job should stay on the start site.
// Image, stylesheet, script, and manifest fetches are never host-bound, so
// CDN files this site publishes are still collected.
func (j job) hostBound() bool {
	switch j.kind {
	case jobImage, jobStyle, jobScript, jobManifest:
		return false
	}
	if looksLikeImage(j.url) {
		return false
	}
	return true
}

func defaultConfig() Config {
	return Config{
		MaxDepth:      5,
		MaxPages:      800,
		SameHost:      true,
		ParseSitemap:  true,
		ParseCSS:      true,
		ParseJS:       true,
		ParseComments: true,
		SaveDataURI:   true,
		RespectRobots: true,
		Concurrency:   6,
		DelayMs:       80,
		MaxBytes:      80 << 20,
	}
}

func Run(ctx context.Context, cfg Config, log LogFn) (*Stats, error) {
	if log == nil {
		log = func(string, string) {}
	}
	start, err := url.Parse(strings.TrimSpace(cfg.StartURL))
	if err != nil || start.Scheme == "" || start.Host == "" {
		return nil, fmt.Errorf("invalid URL: need a full address like https://example.com")
	}
	if start.Scheme != "http" && start.Scheme != "https" {
		return nil, fmt.Errorf("only http and https URLs are supported")
	}
	preview := cfg.Sink != nil
	if !preview {
		if cfg.OutDir == "" {
			cfg.OutDir = "grabbed-images"
		}
		if err := os.MkdirAll(cfg.OutDir, 0o755); err != nil {
			return nil, err
		}
	}
	if cfg.Concurrency < 1 {
		cfg.Concurrency = 4
	}
	if cfg.MaxDepth < 0 {
		cfg.MaxDepth = 0
	}
	if cfg.MaxPages < 1 {
		cfg.MaxPages = 200
	}
	if cfg.MaxBytes < 1 {
		cfg.MaxBytes = 80 << 20
	}

	client := newCrawlClient()
	if u, err := url.Parse(start.Scheme + "://" + start.Host + "/"); err == nil {
		setCookieHeader(client.Jar, u, cfg.Cookies)
	}
	robots := newRobotsCache(client)
	stats := &Stats{}
	originHost := canonicalHost(start)
	homeURL := start.Scheme + "://" + start.Host + "/"

	var (
		mu        sync.Mutex
		visited   = map[string]bool{}
		savedPath = map[string]string{}
		manifest  []ManifestEntry
		queue     = newJobQueue()
		wg        sync.WaitGroup
		pagesLeft = int64(cfg.MaxPages)
		delay     = time.Duration(cfg.DelayMs) * time.Millisecond
		nextID    atomic.Int64
	)

	enqueue := func(j job) {
		if j.url == "" && len(j.data) == 0 {
			return
		}
		u, err := url.Parse(j.url)
		if err != nil && len(j.data) == 0 {
			return
		}
		var unmark func()
		if u != nil {
			if cfg.SameHost && j.hostBound() && u.Scheme != "data" && canonicalHost(u) != originHost {
				return
			}
			if cfg.RespectRobots && j.hostBound() && u.Scheme != "data" && !robots.allowed(u) {
				stats.Skipped.Add(1)
				log("skip", "robots.txt blocked "+j.url)
				return
			}
			key := normalizeVisit(u, cfg.SameHost)
			mu.Lock()
			if visited[key] {
				mu.Unlock()
				return
			}
			visited[key] = true
			mu.Unlock()
			unmark = func() {
				mu.Lock()
				delete(visited, key)
				mu.Unlock()
			}
		}
		wg.Add(1)
		if !queue.Push(ctx, j, cfg.neverDropQueue) {
			if unmark != nil {
				unmark()
			}
			wg.Done()
			if ctx.Err() == nil && !cfg.neverDropQueue() {
				log("warn", "queue full, dropped "+j.url)
			}
		}
	}

	saveBytes := func(hit ImageHit, body []byte, contentType string) {
		if len(body) == 0 {
			return
		}
		name := suggestedName(hit, body, contentType)
		if preview {
			if hit.Hidden {
				stats.Hidden.Add(1)
			}
			id := int(nextID.Add(1))
			img := CollectedImage{
				ID:          id,
				URL:         hit.URL,
				Page:        hit.Page,
				Via:         hit.Via,
				Hidden:      hit.Hidden,
				Bytes:       len(body),
				ContentType: contentType,
				Name:        filepath.ToSlash(name),
				Preview:     fmt.Sprintf("/api/image/%d", id),
				Data:        body,
			}
			cfg.Sink(img)
			mark := ""
			if hit.Hidden {
				mark = " [hidden]"
			}
			log("found", filepath.Base(name)+" ← "+hit.Via+mark)
			return
		}
		written, err := writeImageFile(cfg.OutDir, hit, body, contentType)
		if err != nil {
			stats.Errors.Add(1)
			log("error", err.Error())
			return
		}
		stats.Saved.Add(1)
		if hit.Hidden {
			stats.Hidden.Add(1)
		}
		mu.Lock()
		savedPath[hit.URL] = written
		manifest = append(manifest, ManifestEntry{
			File:   written,
			URL:    hit.URL,
			Page:   hit.Page,
			Via:    hit.Via,
			Hidden: hit.Hidden,
			Bytes:  len(body),
		})
		mu.Unlock()
		mark := ""
		if hit.Hidden {
			mark = " [hidden]"
		}
		log("save", filepath.Base(written)+" ← "+hit.Via+mark)
	}

	noteMiss := func(j job, err error, errLine string) {
		stats.Skipped.Add(1)
		if isQuietMiss(err, j.kind, j.via) {
			return
		}
		log("skip", errLine)
	}

	fetch := func(raw, referer, dest string) ([]byte, string, *url.URL, error) {
		if delay > 0 {
			time.Sleep(delay)
		}
		body, ct, final, err := getURLForgiving(ctx, client, raw, referer, dest, homeURL, cfg.MaxBytes)
		return body, ct, final, err
	}

	handleImageHit := func(img ImageHit, depth int) {
		stats.Found.Add(1)
		if len(img.Data) > 0 {
			if !cfg.SaveDataURI {
				return
			}
			raw := strings.TrimSpace(string(img.Data))
			decoded, err := base64.StdEncoding.DecodeString(raw)
			if err != nil {
				decoded, err = base64.RawStdEncoding.DecodeString(raw)
			}
			if err != nil {
				stats.Errors.Add(1)
				return
			}
			img.URL = "data-uri-" + shortHash(decoded)
			ct := "image/" + strings.TrimPrefix(img.Ext, ".")
			if isVideoName(img.Ext, "", "") {
				ct = "video/" + strings.TrimPrefix(img.Ext, ".")
			}
			saveBytes(img, decoded, ct)
			return
		}
		enqueue(job{kind: jobImage, url: img.URL, depth: depth, via: img.Via, hidden: img.Hidden, page: img.Page})
	}

	process := func(j job) {
		if j.kind == jobImage {
			body, ct, final, err := fetch(j.url, j.page, "media")
			if err != nil {
				noteMiss(j, err, "download "+j.url+": "+err.Error())
				return
			}
			if !isMediaContent(ct, body) && !looksLikeImage(j.url) {
				stats.Skipped.Add(1)
				log("skip", "not media: "+j.url)
				return
			}
			hit := ImageHit{URL: j.url, Page: j.page, Via: j.via, Hidden: j.hidden}
			if final != nil {
				hit.URL = final.String()
			}
			saveBytes(hit, body, ct)
			return
		}

		if j.kind == jobPage {
			n := atomic.AddInt64(&pagesLeft, -1)
			if n < 0 {
				atomic.AddInt64(&pagesLeft, 1)
				return
			}
		}

		body, ct, final, err := fetch(j.url, j.page, "document")
		if err != nil {
			if j.kind == jobPage {
				atomic.AddInt64(&pagesLeft, 1)
			}
			noteMiss(j, err, j.url+": "+err.Error())
			return
		}
		pageURL := final
		if pageURL == nil {
			pageURL, _ = url.Parse(j.url)
		}
		if isMediaContent(ct, body) {
			stats.Found.Add(1)
			saveBytes(ImageHit{URL: pageURL.String(), Page: j.page, Via: j.via, Hidden: j.hidden}, body, ct)
			return
		}

		switch j.kind {
		case jobPage:
			stats.Pages.Add(1)
			log("page", fmt.Sprintf("d%d  %s", j.depth, pageURL))
			if strings.Contains(ct, "json") || strings.HasSuffix(strings.ToLower(pageURL.Path), ".json") {
				for _, img := range extractJSON(pageURL, body) {
					handleImageHit(img, j.depth)
				}
				return
			}
			if strings.Contains(ct, "xml") || strings.HasSuffix(strings.ToLower(pageURL.Path), ".xml") {
				res := extractXML(pageURL, body)
				for _, img := range res.Images {
					handleImageHit(img, j.depth)
				}
				if cfg.ParseSitemap {
					for _, s := range res.Sitemaps {
						enqueue(job{kind: jobSitemap, url: s.URL, depth: j.depth, via: s.Via, page: pageURL.String()})
					}
					for _, p := range res.Pages {
						if j.depth < cfg.MaxDepth {
							enqueue(job{kind: jobPage, url: p.URL, depth: j.depth + 1, via: p.Via, page: pageURL.String()})
						}
					}
				}
				return
			}
			res := extractHTML(pageURL, body)
			for _, img := range res.Images {
				if !cfg.ParseComments && strings.Contains(img.Via, "comment") {
					continue
				}
				handleImageHit(img, j.depth)
			}
			if cfg.ParseCSS {
				for _, s := range res.Styles {
					enqueue(job{kind: jobStyle, url: s.URL, depth: j.depth, via: s.Via, page: pageURL.String()})
				}
			}
			if cfg.ParseJS {
				for _, s := range res.Scripts {
					enqueue(job{kind: jobScript, url: s.URL, depth: j.depth, via: s.Via, page: pageURL.String()})
				}
			}
			for _, s := range res.Manifests {
				enqueue(job{kind: jobManifest, url: s.URL, depth: j.depth, via: s.Via, page: pageURL.String()})
			}
			if cfg.ParseSitemap {
				for _, s := range res.Sitemaps {
					enqueue(job{kind: jobSitemap, url: s.URL, depth: j.depth, via: s.Via, page: pageURL.String()})
				}
			}
			if j.depth < cfg.MaxDepth {
				for _, p := range res.Pages {
					enqueue(job{kind: jobPage, url: p.URL, depth: j.depth + 1, via: p.Via, page: pageURL.String()})
				}
			}
		case jobStyle:
			if !cfg.ParseCSS {
				return
			}
			log("scan", "css "+pageURL.String())
			for _, img := range extractCSS(pageURL, body) {
				handleImageHit(img, j.depth)
			}
		case jobScript:
			if !cfg.ParseJS {
				return
			}
			log("scan", "js "+pageURL.String())
			for _, img := range extractJS(pageURL, body) {
				handleImageHit(img, j.depth)
			}
		case jobSitemap:
			if !cfg.ParseSitemap {
				return
			}
			log("scan", "sitemap "+pageURL.String())
			res := extractXML(pageURL, body)
			for _, img := range res.Images {
				handleImageHit(img, j.depth)
			}
			for _, s := range res.Sitemaps {
				enqueue(job{kind: jobSitemap, url: s.URL, depth: j.depth, via: s.Via, page: pageURL.String()})
			}
			for _, p := range res.Pages {
				if j.depth < cfg.MaxDepth {
					enqueue(job{kind: jobPage, url: p.URL, depth: j.depth + 1, via: p.Via, page: pageURL.String()})
				}
			}
		case jobManifest:
			log("scan", "manifest "+pageURL.String())
			for _, img := range extractJSON(pageURL, body) {
				handleImageHit(img, j.depth)
			}
		}
	}

	var workers sync.WaitGroup
	workers.Add(cfg.Concurrency)
	for i := 0; i < cfg.Concurrency; i++ {
		go func() {
			defer workers.Done()
			for {
				j, ok := queue.Pop()
				if !ok {
					return
				}
				process(j)
				wg.Done()
			}
		}()
	}

	if cfg.ParseSitemap {
		robotsURL, _ := url.Parse(start.Scheme + "://" + start.Host + "/")
		if rf := robots.forURL(robotsURL); rf != nil {
			for _, sm := range rf.sitemaps {
				enqueue(job{kind: jobSitemap, url: sm, depth: 0, via: "robots-sitemap", page: start.String()})
			}
		}
		enqueue(job{kind: jobSitemap, url: start.Scheme + "://" + start.Host + "/sitemap.xml", depth: 0, via: "default-sitemap", page: start.String()})
	}

	if cfg.Cookies != "" {
		log("info", "using session cookies")
	}
	doLogin(ctx, client, cfg, log)
	log("info", "starting at "+start.String()+" — collecting every image this site publishes")
	enqueue(job{kind: jobPage, url: start.String(), depth: 0, via: "start", page: start.String()})
	if dir := listingURLFromStart(start); dir != "" && dir != start.String() {
		enqueue(job{kind: jobPage, url: dir, depth: 0, via: "page-index", page: start.String()})
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		queue.Close()
		workers.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-ctx.Done():
		<-done
	}

	if !preview {
		mu.Lock()
		man := manifest
		mu.Unlock()
		if data, err := json.MarshalIndent(man, "", "  "); err == nil {
			_ = os.WriteFile(filepath.Join(cfg.OutDir, "_manifest.json"), data, 0o644)
		}
	}
	if preview {
		log("info", fmt.Sprintf("done — %d pages, %d found (%d hidden), %d errors — not saved yet",
			stats.Pages.Load(), stats.Found.Load(), stats.Hidden.Load(), stats.Errors.Load()))
	} else {
		log("info", fmt.Sprintf("done — %d pages, %d found, %d saved (%d hidden), %d errors",
			stats.Pages.Load(), stats.Found.Load(), stats.Saved.Load(), stats.Hidden.Load(), stats.Errors.Load()))
	}
	return stats, nil
}

func canonicalHost(u *url.URL) string {
	h := strings.ToLower(u.Hostname())
	h = strings.TrimPrefix(h, "www.")
	port := u.Port()
	if port == "" || (u.Scheme == "http" && port == "80") || (u.Scheme == "https" && port == "443") {
		return h
	}
	return h + ":" + port
}

func normalizeVisit(u *url.URL, sameHost bool) string {
	c := *u
	c.Fragment = ""
	host := strings.ToLower(c.Hostname())
	if sameHost {
		host = strings.TrimPrefix(host, "www.")
	}
	port := c.Port()
	if port == "" || (c.Scheme == "http" && port == "80") || (c.Scheme == "https" && port == "443") {
		c.Host = host
	} else {
		c.Host = host + ":" + port
	}
	return c.String()
}

func shortHash(b []byte) string {
	s := sha1.Sum(b)
	return hex.EncodeToString(s[:8])
}
