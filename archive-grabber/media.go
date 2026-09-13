package main

import (
	"net/url"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// mediaExtAlt is every image and video suffix the grabber will collect from
// markup, CSS, JS, JSON, sitemaps, and directory indexes.
const mediaExtAlt = `jpe?g|jpe|jfif|pjpeg|pjp|png|apng|gif|webp|svgz?|bmp|dib|ico|cur|` +
	`avif|avifs|tiff?|heic|heif|heics|heifs|jxl|jp2|j2k|jpx|jpf|jpm|jxr|wdp|hdp|` +
	`psd|psb|tga|targa|pcx|pbm|pgm|ppm|pnm|xbm|xpm|wbmp|exr|hdr|pic|ras|sgi|` +
	`dng|cr2|cr3|nef|arw|orf|rw2|raf|sr2|srw|kdc|pef|` +
	`mp4|m4v|m4p|mov|qt|webm|ogv|avi|mkv|3gp|3g2|mpe?g|mpe|m1v|m2v|flv|f4v|wmv|asf`

var (
	reMediaExt  = regexp.MustCompile(`(?i)\.(?:` + mediaExtAlt + `)(?:\?|#|$)`)
	reQuotedImg = regexp.MustCompile(`(?i)['"]([^'"]+\.(?:` + mediaExtAlt + `)(?:\?[^'"]*)?)['"]`)
	reBareImg   = regexp.MustCompile(`(?i)(?:https?:)?//[^\s"'<>\\]+\.(?:` + mediaExtAlt + `)(?:\?[^\s"'<>\\]*)?|/[^\s"'<>\\]*\.(?:` + mediaExtAlt + `)(?:\?[^\s"'<>\\]*)?`)
	reDataURI   = regexp.MustCompile(`(?i)data:(image|video)/([a-z0-9.+-]+);base64,([A-Za-z0-9+/=\s]+)`)
)

func looksLikeImage(raw string) bool {
	s := strings.TrimSpace(strings.ToLower(raw))
	if s == "" {
		return false
	}
	if strings.HasPrefix(s, "data:image/") || strings.HasPrefix(s, "data:video/") {
		return true
	}
	u, err := url.Parse(s)
	if err != nil {
		return reMediaExt.MatchString(s)
	}
	return reMediaExt.MatchString(u.Path)
}

func isVideoName(name, rawURL, contentType string) bool {
	ct := strings.ToLower(contentType)
	if strings.HasPrefix(ct, "video/") {
		return true
	}
	ext := strings.ToLower(filepath.Ext(name))
	if ext == "" {
		if u, err := url.Parse(rawURL); err == nil {
			ext = strings.ToLower(path.Ext(u.Path))
		}
	}
	switch ext {
	case ".mp4", ".m4v", ".m4p", ".mov", ".qt", ".webm", ".ogv", ".avi", ".mkv",
		".3gp", ".3g2", ".mpeg", ".mpg", ".mpe", ".m1v", ".m2v", ".flv", ".f4v",
		".wmv", ".asf":
		return true
	}
	return false
}

func isMediaContent(ct string, body []byte) bool {
	ct = strings.ToLower(ct)
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = strings.TrimSpace(ct[:i])
	}
	if strings.HasPrefix(ct, "image/") || strings.HasPrefix(ct, "video/") {
		return true
	}
	if len(body) >= 8 && string(body[:8]) == "\x89PNG\r\n\x1a\n" {
		return true
	}
	if len(body) >= 3 && body[0] == 0xff && body[1] == 0xd8 && body[2] == 0xff {
		return true
	}
	if len(body) >= 6 && (string(body[:6]) == "GIF87a" || string(body[:6]) == "GIF89a") {
		return true
	}
	if len(body) >= 12 && string(body[:4]) == "RIFF" && string(body[8:12]) == "WEBP" {
		return true
	}
	if len(body) >= 12 && string(body[:4]) == "RIFF" && string(body[8:12]) == "AVI " {
		return true
	}
	if len(body) >= 2 && string(body[:2]) == "BM" {
		return true
	}
	if len(body) >= 12 && string(body[4:8]) == "ftyp" {
		return true
	}
	if len(body) >= 4 && body[0] == 0x1a && body[1] == 0x45 && body[2] == 0xdf && body[3] == 0xa3 {
		return true
	}
	if len(body) == 0 {
		return false
	}
	trim := strings.TrimSpace(string(body[:min(256, len(body))]))
	return strings.HasPrefix(trim, "<svg") || (strings.HasPrefix(trim, "<?xml") && strings.Contains(strings.ToLower(trim), "svg"))
}

func mimeToExt(mime string) string {
	switch strings.ToLower(strings.TrimSpace(mime)) {
	case "jpeg", "jpg", "pjpeg":
		return ".jpg"
	case "png", "x-png":
		return ".png"
	case "gif":
		return ".gif"
	case "webp":
		return ".webp"
	case "svg+xml", "svg":
		return ".svg"
	case "bmp", "x-ms-bmp":
		return ".bmp"
	case "x-icon", "vnd.microsoft.icon", "icon":
		return ".ico"
	case "avif":
		return ".avif"
	case "tiff":
		return ".tiff"
	case "apng":
		return ".apng"
	case "heic", "heif":
		return ".heic"
	case "jxl":
		return ".jxl"
	case "mp4", "mpeg4":
		return ".mp4"
	case "webm":
		return ".webm"
	case "quicktime":
		return ".mov"
	case "x-m4v":
		return ".m4v"
	case "ogg", "ogv":
		return ".ogv"
	case "x-msvideo":
		return ".avi"
	case "x-matroska":
		return ".mkv"
	default:
		if strings.Contains(mime, "svg") {
			return ".svg"
		}
		if strings.Contains(mime, "mp4") {
			return ".mp4"
		}
		if strings.HasPrefix(mime, "video/") {
			return ".mp4"
		}
		return ".img"
	}
}

func extToMIME(ext string) string {
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg", ".jpe", ".jfif", ".pjpeg", ".pjp":
		return "image/jpeg"
	case ".png", ".apng":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".svg", ".svgz":
		return "image/svg+xml"
	case ".bmp", ".dib":
		return "image/bmp"
	case ".ico", ".cur":
		return "image/x-icon"
	case ".avif", ".avifs":
		return "image/avif"
	case ".tif", ".tiff":
		return "image/tiff"
	case ".heic", ".heif", ".heics", ".heifs":
		return "image/heic"
	case ".jxl":
		return "image/jxl"
	case ".mp4", ".m4v", ".m4p":
		return "video/mp4"
	case ".webm":
		return "video/webm"
	case ".mov", ".qt":
		return "video/quicktime"
	case ".ogv":
		return "video/ogg"
	case ".avi":
		return "video/x-msvideo"
	case ".mkv":
		return "video/x-matroska"
	case ".mpeg", ".mpg", ".mpe", ".m1v", ".m2v":
		return "video/mpeg"
	case ".3gp":
		return "video/3gpp"
	case ".3g2":
		return "video/3gpp2"
	case ".flv", ".f4v":
		return "video/x-flv"
	case ".wmv", ".asf":
		return "video/x-ms-wmv"
	default:
		return ""
	}
}

func previewContentType(ct, name, rawURL string) string {
	ct = strings.ToLower(strings.TrimSpace(ct))
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = strings.TrimSpace(ct[:i])
	}
	if strings.HasPrefix(ct, "image/") || strings.HasPrefix(ct, "video/") {
		return ct
	}
	ext := strings.ToLower(filepath.Ext(name))
	if ext == "" {
		if u, err := url.Parse(rawURL); err == nil {
			ext = strings.ToLower(path.Ext(u.Path))
		}
	}
	if mime := extToMIME(ext); mime != "" {
		return mime
	}
	return "application/octet-stream"
}
