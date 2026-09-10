package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

const browserUA = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36"

var errForbidden = errors.New("HTTP 403")

var userFolders = []string{
	"users", "user", "members", "member", "profiles", "profile",
	"people", "avatars", "avatar", "accounts", "u",
}

var mediaFolders = []string{
	"images", "img", "image", "photos", "photo", "pics", "pictures",
	"media", "uploads", "upload", "files", "gallery", "galleries",
	"albums", "thumbs", "thumbnails", "static", "assets", "content",
	"videos", "video", "clips", "movies", "mp4", "footage",
}

var extraAdminFolders = []string{
	"attachments", "storage", "public", "data", "download", "downloads",
	"wp-content/uploads", "media/users", "user/uploads", "users/uploads",
	"users/images", "profile/photos", "profiles/photos", "files/users",
}

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

func isLoginPath(p string) bool {
	p = strings.ToLower(p)
	for _, n := range []string{"/login", "/signin", "/sign-in", "/log-in", "/account/login", "/users/sign_in", "/auth/login"} {
		if strings.Contains(p, n) {
			return true
		}
	}
	return false
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
	if strings.Contains(msg, "HTTP 404") && (via == "folder-seed" || via == "users-seed" || via == "default-sitemap" || via == "robots-sitemap") {
		return true
	}
	return false
}

func folderSeeds(base *url.URL, deep bool) []job {
	if base == nil {
		return nil
	}
	names := append([]string{}, userFolders...)
	names = append(names, mediaFolders...)
	if deep {
		names = append(names, extraAdminFolders...)
	}
	seen := map[string]bool{}
	var out []job
	root := *base
	root.Path = "/"
	root.RawQuery = ""
	root.Fragment = ""
	for _, name := range names {
		p := "/" + strings.Trim(name, "/") + "/"
		if seen[p] {
			continue
		}
		seen[p] = true
		u := root
		u.Path = p
		via := "folder-seed"
		if strings.Contains(p, "user") || strings.Contains(p, "member") || strings.Contains(p, "profile") || strings.Contains(p, "avatar") || p == "/u/" || strings.Contains(p, "people") || strings.Contains(p, "account") {
			via = "users-seed"
		}
		out = append(out, job{kind: jobPage, url: u.String(), depth: 0, via: via, page: base.String()})
	}
	return out
}
