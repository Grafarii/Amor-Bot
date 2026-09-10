package main

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

const defaultSafeQueueCap = 8192

var safeQueueCap = defaultSafeQueueCap

// queueOfferWait is how long safe mode waits for a slot before dropping.
// Tests may set this to 0 for a deterministic immediate drop.
var queueOfferWait = 2 * time.Second

type jobQueue struct {
	mu     sync.Mutex
	cond   *sync.Cond
	items  []job
	head   int
	closed bool
}

func newJobQueue() *jobQueue {
	q := &jobQueue{items: make([]job, 0, 256)}
	q.cond = sync.NewCond(&q.mu)
	return q
}

func (q *jobQueue) lenLocked() int {
	return len(q.items) - q.head
}

func (q *jobQueue) compactLocked() {
	if q.head > 1024 && q.head*2 >= len(q.items) {
		q.items = append([]job(nil), q.items[q.head:]...)
		q.head = 0
	}
}

func (q *jobQueue) Push(ctx context.Context, j job, neverDrop func() bool) bool {
	if neverDrop == nil {
		neverDrop = func() bool { return false }
	}
	if ctx.Err() != nil {
		return false
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	for {
		if q.closed || ctx.Err() != nil {
			return false
		}
		if neverDrop() || q.lenLocked() < safeQueueCap {
			q.items = append(q.items, j)
			q.cond.Signal()
			return true
		}
		if queueOfferWait <= 0 {
			return false
		}

		var timedOut atomic.Bool
		timer := time.AfterFunc(queueOfferWait, func() {
			timedOut.Store(true)
			q.cond.Broadcast()
		})
		stop := context.AfterFunc(ctx, func() {
			q.cond.Broadcast()
		})
		q.cond.Wait()
		timer.Stop()
		stop()

		if q.closed || ctx.Err() != nil {
			return false
		}
		if neverDrop() || q.lenLocked() < safeQueueCap {
			q.items = append(q.items, j)
			q.cond.Signal()
			return true
		}
		if timedOut.Load() {
			return false
		}
	}
}

func (q *jobQueue) Pop() (job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for q.lenLocked() == 0 && !q.closed {
		q.cond.Wait()
	}
	if q.lenLocked() == 0 {
		return job{}, false
	}
	j := q.items[q.head]
	q.items[q.head] = job{}
	q.head++
	q.compactLocked()
	q.cond.Signal()
	return j, true
}

func (q *jobQueue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.lenLocked()
}

func (q *jobQueue) Close() {
	q.mu.Lock()
	q.closed = true
	q.cond.Broadcast()
	q.mu.Unlock()
}
