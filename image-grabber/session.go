package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"
)

const browserUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"

var errForbidden = errors.New("HTTP 403")

func applyBrowserHeaders(req *http.Request, referer, dest string) {
	req.Header.Set("User-Agent", browserUA)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	switch dest {
	case "image", "media":
		req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/*,video/webm,video/mp4,video/*,*/*;q=0.8")
		if dest == "media" {
			req.Header.Set("Sec-Fetch-Dest", "empty")
		} else {
			req.Header.Set("Sec-Fetch-Dest", "image")
		}
		req.Header.Set("Sec-Fetch-Mode", "no-cors")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
	default:
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
	}
	if referer != "" {
		req.Header.Set("Referer", referer)
	}
}

func newCrawlClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	return &http.Client{
		Timeout: 25 * time.Second,
		Jar:     jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 8 {
				return errors.New("too many redirects")
			}
			return nil
		},
	}
}

func setCookieHeader(jar http.CookieJar, page *url.URL, header string) {
	if jar == nil || page == nil || strings.TrimSpace(header) == "" {
		return
	}
	var cookies []*http.Cookie
	for _, part := range strings.Split(header, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		name, val, ok := strings.Cut(part, "=")
		name = strings.TrimSpace(name)
		val = strings.TrimSpace(val)
		if !ok || name == "" {
			continue
		}
		cookies = append(cookies, &http.Cookie{Name: name, Value: val, Path: "/"})
	}
	if len(cookies) > 0 {
		jar.SetCookies(page, cookies)
	}
}

func doLogin(ctx context.Context, client *http.Client, cfg Config, log LogFn) {
	if strings.TrimSpace(cfg.LoginURL) == "" || strings.TrimSpace(cfg.LoginUser) == "" {
		return
	}
	form := url.Values{}
	form.Set("username", cfg.LoginUser)
	form.Set("user", cfg.LoginUser)
	form.Set("email", cfg.LoginUser)
	form.Set("login", cfg.LoginUser)
	form.Set("password", cfg.LoginPass)
	form.Set("pass", cfg.LoginPass)
	form.Set("passwd", cfg.LoginPass)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, cfg.LoginURL, strings.NewReader(form.Encode()))
	if err != nil {
		log("warn", "login request: "+err.Error())
		return
	}
	applyBrowserHeaders(req, cfg.LoginURL, "document")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", originOf(cfg.LoginURL))
	resp, err := client.Do(req)
	if err != nil {
		log("warn", "login failed: "+err.Error())
		return
	}
	io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<20))
	resp.Body.Close()
	if resp.StatusCode >= 400 && resp.StatusCode != 302 {
		log("warn", "login HTTP "+resp.Status)
		return
	}
	log("info", "session login posted — continuing with cookies")
}

func originOf(raw string) string {
	u, err := url.Parse(raw)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ""
	}
	return u.Scheme + "://" + u.Host
}

// listingURLFromStart returns the directory index next to a start URL like
// /gallery/index.html, so the crawl sniffs the listing itself.
func listingURLFromStart(start *url.URL) string {
	if start == nil {
		return ""
	}
	base := strings.ToLower(path.Base(start.Path))
	switch base {
	case "index.html", "index.htm", "index.php", "default.html", "default.htm":
	default:
		return ""
	}
	u := *start
	dir := path.Dir(start.Path)
	if dir == "." || dir == "/" {
		u.Path = "/"
	} else {
		u.Path = strings.TrimSuffix(dir, "/") + "/"
	}
	u.RawQuery = ""
	u.Fragment = ""
	return u.String()
}

func isSoftMiss(err error, via string) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, errForbidden) {
		return true
	}
	msg := err.Error()
	if strings.Contains(msg, "HTTP 403") || strings.Contains(msg, "HTTP 401") {
		return true
	}
	return strings.Contains(msg, "HTTP 404") && (via == "default-sitemap" || via == "robots-sitemap")
}
