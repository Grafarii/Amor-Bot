package main

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type robotsGroup struct {
	disallow []string
	allow    []string
}

type robotsFile struct {
	groups    []robotsGroup
	starIndex int
	sitemaps  []string
}

type robotsCache struct {
	mu     sync.Mutex
	byHost map[string]*robotsFile
	client *http.Client
}

func newRobotsCache(client *http.Client) *robotsCache {
	return &robotsCache{byHost: map[string]*robotsFile{}, client: client}
}

func (c *robotsCache) forURL(u *url.URL) *robotsFile {
	host := strings.ToLower(u.Hostname())
	c.mu.Lock()
	if f, ok := c.byHost[host]; ok {
		c.mu.Unlock()
		return f
	}
	c.mu.Unlock()

	robotsURL := u.Scheme + "://" + u.Host + "/robots.txt"
	f := &robotsFile{}
	req, err := http.NewRequest(http.MethodGet, robotsURL, nil)
	if err == nil {
		req.Header.Set("User-Agent", userAgent)
		resp, err := c.client.Do(req)
		if err == nil {
			defer resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				f = parseRobots(resp.Body)
			}
		}
	}

	c.mu.Lock()
	c.byHost[host] = f
	c.mu.Unlock()
	return f
}

func (c *robotsCache) allowed(u *url.URL) bool {
	f := c.forURL(u)
	if f == nil {
		return true
	}
	if f.starIndex < 0 || f.starIndex >= len(f.groups) {
		return true
	}
	g := &f.groups[f.starIndex]
	path := u.EscapedPath()
	if path == "" {
		path = "/"
	}
	allowLen, disallowLen := -1, -1
	for _, p := range g.allow {
		if strings.HasPrefix(path, p) && len(p) > allowLen {
			allowLen = len(p)
		}
	}
	for _, p := range g.disallow {
		if p == "" {
			continue
		}
		if strings.HasPrefix(path, p) && len(p) > disallowLen {
			disallowLen = len(p)
		}
	}
	if allowLen == -1 && disallowLen == -1 {
		return true
	}
	return allowLen >= disallowLen
}

func parseRobots(r io.Reader) *robotsFile {
	f := &robotsFile{starIndex: -1}
	current := -1
	applies := false
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if i := strings.Index(line, "#"); i >= 0 {
			line = strings.TrimSpace(line[:i])
		}
		if line == "" {
			continue
		}
		key, val, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		key = strings.ToLower(strings.TrimSpace(key))
		val = strings.TrimSpace(val)
		switch key {
		case "user-agent":
			f.groups = append(f.groups, robotsGroup{})
			current = len(f.groups) - 1
			applies = val == "*" || strings.Contains(strings.ToLower(userAgent), strings.ToLower(val))
			if val == "*" {
				f.starIndex = current
			}
		case "disallow":
			if applies && current >= 0 {
				f.groups[current].disallow = append(f.groups[current].disallow, val)
			}
		case "allow":
			if applies && current >= 0 {
				f.groups[current].allow = append(f.groups[current].allow, val)
			}
		case "sitemap":
			if val != "" {
				f.sitemaps = append(f.sitemaps, val)
			}
		}
	}
	if f.starIndex < 0 && len(f.groups) > 0 {
		f.starIndex = 0
	}
	return f
}
