package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

var winIllegal = regexp.MustCompile(`[<>:"|?*]`)

func writeImageFile(outDir string, hit ImageHit, body []byte, contentType string) (string, error) {
	rel := suggestedName(hit, body, contentType)
	full := filepath.Join(outDir, rel)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return "", err
	}
	f, full, err := createUnique(full)
	if err != nil {
		return "", err
	}
	_, err = f.Write(body)
	if err != nil {
		f.Close()
		os.Remove(full)
		return "", err
	}
	if err := f.Close(); err != nil {
		os.Remove(full)
		return "", err
	}
	rel, err = filepath.Rel(outDir, full)
	if err != nil {
		rel = filepath.Base(full)
	}
	return rel, nil
}

func suggestedName(hit ImageHit, body []byte, contentType string) string {
	ext := extFrom(hit, body, contentType)
	if len(hit.Data) > 0 || strings.HasPrefix(hit.URL, "data") {
		sum := sha1.Sum(body)
		return filepath.Join("_inline", hex.EncodeToString(sum[:8])+ext)
	}
	u, err := url.Parse(hit.URL)
	if err != nil {
		sum := sha1.Sum(body)
		return hex.EncodeToString(sum[:10]) + ext
	}
	host := sanitizeSegment(u.Hostname())
	if host == "" {
		host = "site"
	}
	p := u.Path
	if strings.HasSuffix(p, "/") || p == "" {
		p += "index" + ext
	}
	p = strings.TrimPrefix(p, "/")
	parts := strings.Split(p, "/")
	clean := make([]string, 0, len(parts)+1)
	clean = append(clean, host)
	for _, part := range parts {
		part = sanitizeSegment(part)
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if len(clean) > 1 {
				clean = clean[:len(clean)-1]
			}
			continue
		}
		clean = append(clean, part)
	}
	name := filepath.Join(clean...)
	if filepath.Ext(name) == "" {
		name += ext
	} else if ext != "" && strings.ToLower(filepath.Ext(name)) != strings.ToLower(ext) && isGenericExt(filepath.Ext(name)) {
		name += ext
	}
	if u.RawQuery != "" {
		base := strings.TrimSuffix(name, filepath.Ext(name))
		q := sanitizeSegment(u.RawQuery)
		if len(q) > 40 {
			sum := sha1.Sum([]byte(u.RawQuery))
			q = hex.EncodeToString(sum[:6])
		}
		name = base + "_" + q + filepath.Ext(name)
	}
	if len(name) > 180 {
		sum := sha1.Sum([]byte(hit.URL))
		name = filepath.Join(host, hex.EncodeToString(sum[:10])+ext)
	}
	return name
}

func extFrom(hit ImageHit, body []byte, contentType string) string {
	if hit.Ext != "" {
		if !strings.HasPrefix(hit.Ext, ".") {
			return "." + hit.Ext
		}
		return hit.Ext
	}
	ct := strings.ToLower(contentType)
	switch {
	case strings.Contains(ct, "jpeg"), strings.Contains(ct, "jpg"):
		return ".jpg"
	case strings.Contains(ct, "png"):
		return ".png"
	case strings.Contains(ct, "gif"):
		return ".gif"
	case strings.Contains(ct, "webp"):
		return ".webp"
	case strings.Contains(ct, "svg"):
		return ".svg"
	case strings.Contains(ct, "avif"):
		return ".avif"
	case strings.Contains(ct, "bmp"):
		return ".bmp"
	case strings.Contains(ct, "icon"):
		return ".ico"
	case strings.Contains(ct, "tiff"):
		return ".tiff"
	case strings.Contains(ct, "heic"), strings.Contains(ct, "heif"):
		return ".heic"
	case strings.Contains(ct, "jxl"):
		return ".jxl"
	case strings.Contains(ct, "mp4"):
		return ".mp4"
	case strings.Contains(ct, "webm"):
		return ".webm"
	case strings.Contains(ct, "quicktime"):
		return ".mov"
	case strings.Contains(ct, "ogg") && strings.Contains(ct, "video"):
		return ".ogv"
	case strings.Contains(ct, "msvideo") || strings.Contains(ct, "avi"):
		return ".avi"
	case strings.Contains(ct, "matroska"):
		return ".mkv"
	case strings.HasPrefix(ct, "video/"):
		return ".mp4"
	}
	if looksLikeImage(hit.URL) {
		u, err := url.Parse(hit.URL)
		if err == nil {
			ext := strings.ToLower(filepath.Ext(u.Path))
			if ext != "" {
				return ext
			}
		}
	}
	if len(body) >= 12 && string(body[4:8]) == "ftyp" {
		return ".mp4"
	}
	if len(body) >= 8 && string(body[:8]) == "\x89PNG\r\n\x1a\n" {
		return ".png"
	}
	if len(body) >= 3 && body[0] == 0xff && body[1] == 0xd8 {
		return ".jpg"
	}
	if len(body) >= 6 && strings.HasPrefix(string(body), "GIF") {
		return ".gif"
	}
	if len(body) >= 12 && string(body[8:12]) == "WEBP" {
		return ".webp"
	}
	trim := strings.TrimSpace(string(body[:min(120, len(body))]))
	if strings.HasPrefix(trim, "<svg") {
		return ".svg"
	}
	return ".img"
}

func isGenericExt(ext string) bool {
	switch strings.ToLower(ext) {
	case ".php", ".asp", ".aspx", ".jsp", ".html", ".htm", "":
		return true
	}
	return false
}

func sanitizeSegment(s string) string {
	s = strings.TrimSpace(s)
	s = winIllegal.ReplaceAllString(s, "_")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "\\", "_")
	s = strings.Map(func(r rune) rune {
		if unicode.IsControl(r) {
			return -1
		}
		if r == 0 {
			return -1
		}
		return r
	}, s)
	s = strings.Trim(s, " .")
	if s == "" {
		return ""
	}
	reserved := map[string]bool{"con": true, "prn": true, "aux": true, "nul": true}
	if reserved[strings.ToLower(s)] {
		return "_" + s
	}
	return s
}

func createUnique(full string) (*os.File, string, error) {
	ext := filepath.Ext(full)
	base := strings.TrimSuffix(full, ext)
	for i := 1; i < 1000; i++ {
		cand := full
		if i > 1 {
			cand = fmt.Sprintf("%s_%d%s", base, i, ext)
		}
		f, err := os.OpenFile(cand, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
		if err == nil {
			return f, cand, nil
		}
		if !os.IsExist(err) {
			return nil, "", err
		}
	}
	return nil, "", fmt.Errorf("too many name collisions for %s", filepath.Base(full))
}
