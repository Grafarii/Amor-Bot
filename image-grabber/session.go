package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strings"
	"time"
)

const browserUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36"

var errForbidden = errors.New("HTTP 403")

func applyBrowserHeaders(req *http.Request, referer, dest string) {
	applyBrowserHeadersOpts(req, referer, dest, false)
}

func applyBrowserHeadersOpts(req *http.Request, referer, dest string, withOrigin bool) {
	req.Header.Set("User-Agent", browserUA)
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,pt-BR;q=0.8,pt;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="131", "Not_A Brand";v="24", "Google Chrome";v="131"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	site := "none"
	if referer != "" {
		if ru, err := url.Parse(referer); err == nil {
			site = secFetchSite(ru, req.URL)
		}
	}
	media := dest == "image" || dest == "media"
	if media && looksLikeVideoURL(req.URL) {
		req.Header.Set("Accept", "video/webm,video/mp4,video/*,*/*;q=0.8")
		req.Header.Set("Sec-Fetch-Dest", "video")
		req.Header.Set("Sec-Fetch-Mode", "no-cors")
	} else if media {
		req.Header.Set("Accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8")
		req.Header.Set("Sec-Fetch-Dest", "image")
		req.Header.Set("Sec-Fetch-Mode", "no-cors")
	} else {
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-User", "?1")
	}
	req.Header.Set("Sec-Fetch-Site", site)
	if referer != "" {
		req.Header.Set("Referer", referer)
	} else {
		req.Header.Del("Referer")
	}
	if withOrigin {
		if o := originOf(referer); o != "" {
			req.Header.Set("Origin", o)
		}
	} else {
		req.Header.Del("Origin")
	}
}

func looksLikeVideoURL(u *url.URL) bool {
	if u == nil {
		return false
	}
	return isVideoName("", u.String(), "")
}

func secFetchSite(referer, req *url.URL) string {
	if referer == nil || req == nil || referer.Host == "" {
		return "none"
	}
	if strings.EqualFold(referer.Scheme, req.Scheme) && canonicalHost(referer) == canonicalHost(req) {
		return "same-origin"
	}
	rh := strings.TrimPrefix(strings.ToLower(referer.Hostname()), "www.")
	qh := strings.TrimPrefix(strings.ToLower(req.Hostname()), "www.")
	if rh == "" || qh == "" {
		return "cross-site"
	}
	if rh == qh || strings.HasSuffix(qh, "."+rh) || strings.HasSuffix(rh, "."+qh) {
		return "same-site"
	}
	return "cross-site"
}

func newCrawlClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	return &http.Client{
		Timeout: 25 * time.Second,
		Jar:     jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return errors.New("too many redirects")
			}
			prev := via[len(via)-1]
			dest := "document"
			switch prev.Header.Get("Sec-Fetch-Dest") {
			case "image", "video", "empty":
				dest = "image"
			}
			applyBrowserHeaders(req, prev.URL.String(), dest)
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
	code := httpStatusFromErr(err)
	switch code {
	case 401, 403, 404, 405, 406, 408, 410, 429, 451:
		return true
	}
	msg := err.Error()
	if strings.Contains(msg, "HTTP 403") || strings.Contains(msg, "HTTP 401") {
		return true
	}
	return strings.Contains(msg, "HTTP 404") && (via == "default-sitemap" || via == "robots-sitemap")
}

func httpStatusFromErr(err error) int {
	if err == nil {
		return 0
	}
	msg := err.Error()
	var n int
	if _, scanErr := fmt.Sscanf(msg, "HTTP %d", &n); scanErr == nil {
		return n
	}
	return 0
}

func isQuietMiss(err error, kind jobKind, via string) bool {
	if !isSoftMiss(err, via) {
		return false
	}
	code := httpStatusFromErr(err)
	if errors.Is(err, errForbidden) || code == 401 || code == 403 {
		return true
	}
	if code == 404 || strings.Contains(err.Error(), "HTTP 404") {
		switch kind {
		case jobStyle, jobScript, jobSitemap, jobManifest:
			return true
		}
		switch via {
		case "default-sitemap", "robots-sitemap", "stylesheet", "script-src":
			return true
		}
	}
	return false
}

func getURL(ctx context.Context, client *http.Client, raw, referer, dest string, withOrigin bool, maxBytes int64) ([]byte, string, *url.URL, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, raw, nil)
	if err != nil {
		return nil, "", nil, 0, err
	}
	applyBrowserHeadersOpts(req, referer, dest, withOrigin)
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", nil, 0, err
	}
	defer resp.Body.Close()
	final := resp.Request.URL
	if maxBytes <= 0 {
		maxBytes = 80 << 20
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBytes+1))
	if err != nil {
		return nil, "", final, resp.StatusCode, err
	}
	if int64(len(body)) > maxBytes {
		return nil, "", final, resp.StatusCode, fmt.Errorf("response too large")
	}
	ct := resp.Header.Get("Content-Type")
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = ct[:i]
	}
	ct = strings.TrimSpace(strings.ToLower(ct))
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return body, ct, final, resp.StatusCode, nil
	}
	if bodyUsable(resp.StatusCode, ct, body) {
		return body, ct, final, resp.StatusCode, nil
	}
	return nil, ct, final, resp.StatusCode, fmt.Errorf("HTTP %d", resp.StatusCode)
}

func bodyUsable(code int, ct string, body []byte) bool {
	if len(body) == 0 {
		return false
	}
	switch code {
	case 401, 403, 404, 410:
	default:
		return false
	}
	if isMediaContent(ct, body) {
		return true
	}
	c := strings.ToLower(ct)
	if strings.Contains(c, "css") || strings.Contains(c, "javascript") || strings.Contains(c, "json") || strings.Contains(c, "xml") {
		return len(body) > 8
	}
	s := strings.ToLower(string(body))
	if strings.Contains(s, "<img") || strings.Contains(s, "<source") || strings.Contains(s, "<video") ||
		strings.Contains(s, "<picture") || strings.Contains(s, "srcset") || strings.Contains(s, "data-src") {
		return true
	}
	if strings.Contains(c, "html") || strings.Contains(s, "<html") || strings.Contains(s, "<!doctype") {
		if strings.Contains(s, "href=") || strings.Contains(s, "stylesheet") || strings.Contains(s, "<script") {
			return len(body) > 64
		}
	}
	return false
}

func hotlinkReferers(page, raw, home string) []string {
	var out []string
	seen := map[string]bool{}
	add := func(s string) {
		if seen[s] {
			return
		}
		seen[s] = true
		out = append(out, s)
	}
	add(page)
	add(home)
	if u, err := url.Parse(page); err == nil && u.Host != "" {
		add(u.Scheme + "://" + u.Host + "/")
		host := u.Hostname()
		if strings.HasPrefix(strings.ToLower(host), "www.") {
			add(u.Scheme + "://" + host[4:] + "/")
		} else if host != "" {
			add(u.Scheme + "://www." + host + "/")
		}
	}
	add(raw)
	add("")
	return out
}

// getURLForgiving fetches like a browser and retries the usual hotlink 403 cases
// (wrong Referer, missing cookies, cross-site Sec-Fetch-Site). A true denial
// still returns errForbidden so the crawl skips instead of failing.
func getURLForgiving(ctx context.Context, client *http.Client, raw, page, dest, home string, maxBytes int64) ([]byte, string, *url.URL, error) {
	if page == "" {
		page = home
	}
	type attempt struct {
		referer    string
		dest       string
		withOrigin bool
	}
	var tries []attempt
	seen := map[string]bool{}
	add := func(a attempt) {
		key := a.referer + "\x00" + a.dest + "\x00" + fmt.Sprint(a.withOrigin)
		if seen[key] {
			return
		}
		seen[key] = true
		tries = append(tries, a)
	}
	refs := hotlinkReferers(page, raw, home)
	for _, ref := range refs {
		add(attempt{referer: ref, dest: dest})
	}
	if dest != "document" {
		add(attempt{referer: page, dest: "document"})
		add(attempt{referer: page, dest: dest, withOrigin: true})
		add(attempt{referer: "", dest: dest, withOrigin: false})
	} else {
		add(attempt{referer: page, dest: "image"})
		add(attempt{referer: "", dest: "document"})
	}

	var lastFinal *url.URL
	var lastErr error
	var lastCode int
	for i, a := range tries {
		if i > 0 {
			select {
			case <-ctx.Done():
				return nil, "", lastFinal, ctx.Err()
			case <-time.After(20 * time.Millisecond):
			}
		}
		body, ct, final, code, err := getURL(ctx, client, raw, a.referer, a.dest, a.withOrigin, maxBytes)
		lastFinal, lastErr, lastCode = final, err, code
		if err == nil {
			return body, ct, final, nil
		}
		if code != http.StatusForbidden && code != http.StatusUnauthorized {
			return body, ct, final, err
		}
		// Same headers once more so a Set-Cookie from the 403 can pass.
		if i == 0 {
			body, ct, final, code, err = getURL(ctx, client, raw, a.referer, a.dest, a.withOrigin, maxBytes)
			lastFinal, lastErr, lastCode = final, err, code
			if err == nil {
				return body, ct, final, nil
			}
			if code != http.StatusForbidden && code != http.StatusUnauthorized {
				return body, ct, final, err
			}
		}
		if len(tries) > 8 && i >= 7 {
			break
		}
	}
	if lastCode == http.StatusForbidden || lastCode == http.StatusUnauthorized || (lastErr != nil && isSoftMiss(lastErr, "")) {
		return nil, "", lastFinal, errForbidden
	}
	return nil, "", lastFinal, lastErr
}
