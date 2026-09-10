package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestJobQueueNeverDropsWhenUnlocked(t *testing.T) {
	q := newJobQueue()
	ctx := context.Background()
	oldWait := queueOfferWait
	queueOfferWait = 0
	t.Cleanup(func() { queueOfferWait = oldWait })

	for i := 0; i < safeQueueCap; i++ {
		if !q.Push(ctx, job{url: fmt.Sprintf("http://example.test/%d", i)}, func() bool { return false }) {
			t.Fatalf("safe fill failed at %d", i)
		}
	}
	if q.Push(ctx, job{url: "http://example.test/overflow"}, func() bool { return false }) {
		t.Fatal("safe mode should drop once the queue is full")
	}
	if !q.Push(ctx, job{url: "http://example.test/unlocked"}, func() bool { return true }) {
		t.Fatal("unlocked mode must accept work past the safe cap")
	}
	if q.Len() != safeQueueCap+1 {
		t.Fatalf("len=%d, want %d", q.Len(), safeQueueCap+1)
	}
}

func TestUnlockDuringWaitAcceptsJob(t *testing.T) {
	q := newJobQueue()
	ctx := context.Background()
	oldWait := queueOfferWait
	queueOfferWait = 150 * time.Millisecond
	t.Cleanup(func() { queueOfferWait = oldWait })

	for i := 0; i < safeQueueCap; i++ {
		if !q.Push(ctx, job{url: fmt.Sprintf("http://example.test/%d", i)}, func() bool { return false }) {
			t.Fatalf("fill failed at %d", i)
		}
	}

	var unlocked atomic.Bool
	done := make(chan bool, 1)
	go func() {
		done <- q.Push(ctx, job{url: "http://example.test/late"}, unlocked.Load)
	}()
	time.Sleep(30 * time.Millisecond)
	unlocked.Store(true)
	q.cond.Broadcast()

	select {
	case ok := <-done:
		if !ok {
			t.Fatal("unlock during wait should keep the job")
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for unlocked push")
	}
}

func TestNeverDropQueueFollowsLiveAtomic(t *testing.T) {
	flag := &atomic.Bool{}
	cfg := Config{DeepScan: false, Unlimited: flag}
	if cfg.neverDropQueue() {
		t.Fatal("locked crawl should still drop")
	}
	flag.Store(true)
	if !cfg.neverDropQueue() {
		t.Fatal("unlocking admin should stop drops immediately")
	}
}

func TestLockedCrawlCanLogQueueFull(t *testing.T) {
	oldCap := safeQueueCap
	oldWait := queueOfferWait
	safeQueueCap = 2
	queueOfferWait = 0
	t.Cleanup(func() {
		safeQueueCap = oldCap
		queueOfferWait = oldWait
	})

	var b strings.Builder
	for i := 0; i < 80; i++ {
		fmt.Fprintf(&b, `<img src="/img-%d.png">`, i)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/img-") && strings.HasSuffix(r.URL.Path, ".png") {
			w.Header().Set("Content-Type", "image/png")
			w.Write(pngDot)
			return
		}
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html><body>" + b.String() + "</body></html>"))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	var logs []string
	var mu sync.Mutex
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Concurrency = 1
	cfg.MaxPages = 2
	cfg.DeepScan = false
	cfg.Sink = func(CollectedImage) {}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := Run(ctx, cfg, func(kind, msg string) {
		mu.Lock()
		logs = append(logs, kind+" "+msg)
		mu.Unlock()
	}); err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	defer mu.Unlock()
	for _, line := range logs {
		if strings.Contains(line, "queue full") {
			return
		}
	}
	t.Fatal("safe mode should log queue full when the cap is tiny")
}

func TestUnlockedCrawlDoesNotLogQueueFull(t *testing.T) {
	oldCap := safeQueueCap
	safeQueueCap = 8
	t.Cleanup(func() { safeQueueCap = oldCap })

	var b strings.Builder
	for i := 0; i < 80; i++ {
		fmt.Fprintf(&b, `<img src="/img-%d.png">`, i)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/img-") && strings.HasSuffix(r.URL.Path, ".png") {
			w.Header().Set("Content-Type", "image/png")
			w.Write(pngDot)
			return
		}
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html><body>" + b.String() + "</body></html>"))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	var logs []string
	var mu sync.Mutex
	unlocked := &atomic.Bool{}
	unlocked.Store(true)
	cfg := defaultConfig()
	cfg.StartURL = srv.URL + "/"
	cfg.DelayMs = 0
	cfg.RespectRobots = false
	cfg.ParseSitemap = false
	cfg.Concurrency = 1
	cfg.MaxPages = 2
	cfg.Sink = func(CollectedImage) {}
	cfg.Unlimited = unlocked
	applyBounds(&cfg, true)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	if _, err := Run(ctx, cfg, func(kind, msg string) {
		mu.Lock()
		logs = append(logs, kind+" "+msg)
		mu.Unlock()
	}); err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	defer mu.Unlock()
	for _, line := range logs {
		if strings.Contains(line, "queue full") {
			t.Fatalf("unlocked crawl logged %q", line)
		}
	}
}
