package main

import (
	"bytes"
	"net/url"
	"path"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

var (
	reCSSURL      = regexp.MustCompile(`(?i)url\(\s*['"]?([^'")]+)['"]?\s*\)`)
	reXMLLoc      = regexp.MustCompile(`(?is)<(?:image:)?loc>\s*([^<]+?)\s*</(?:image:)?loc>`)
	reMediaURL    = regexp.MustCompile(`(?i)(?:url|href|src|content)=["']([^"']+)["']`)
	reHiddenClass = regexp.MustCompile(`(?i)(?:^|\s)(?:hidden|hide|invisible|sr-only|visually-hidden|d-none|is-hidden|u-hidden|display-none|visuallyhidden)(?:\s|$)`)
	reDisplayNone = regexp.MustCompile(`(?i)display\s*:\s*none|visibility\s*:\s*hidden|opacity\s*:\s*0(?:\.0+)?(?:\s|;|$)`)
	reIndexFile   = regexp.MustCompile(`(?i)(?:^|[\s"'>=])([A-Za-z0-9._~+-]+\.(?:` + mediaExtAlt + `))`)
)

var voidTags = map[string]bool{
	"area": true, "base": true, "br": true, "col": true, "embed": true,
	"hr": true, "img": true, "input": true, "link": true, "meta": true,
	"param": true, "source": true, "track": true, "wbr": true,
}

var pageExt = map[string]bool{
	"": true, ".html": true, ".htm": true, ".php": true, ".asp": true,
	".aspx": true, ".jsp": true, ".xhtml": true, ".shtml": true,
}

type ImageHit struct {
	URL    string
	Page   string
	Via    string
	Hidden bool
	Data   []byte // set for data: URIs
	Ext    string
}

type ResourceHit struct {
	URL string
	Via string
}

type ExtractResult struct {
	Images    []ImageHit
	Pages     []ResourceHit
	Styles    []ResourceHit
	Scripts   []ResourceHit
	Sitemaps  []ResourceHit
	Manifests []ResourceHit
}

func extractHTML(page *url.URL, body []byte) ExtractResult {
	var out ExtractResult
	seenImg := map[string]bool{}
	indexListing := isDirectoryIndex(page, body)
	addImg := func(raw, via string, hidden bool) {
		addImage(&out, seenImg, page, raw, via, hidden)
	}

	z := html.NewTokenizer(bytes.NewReader(body))
	hiddenStack := []bool{false}
	var (
		inStyle, inScript, inNoscript bool
		scriptType                    string
		textBuf                       strings.Builder
	)

	hiddenNow := func() bool {
		return hiddenStack[len(hiddenStack)-1] || inNoscript
	}

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			if indexListing {
				sniffIndexListing(page, body, &out, seenImg)
			}
			rawPass(page, body, &out, seenImg)
			return out

		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			name := strings.ToLower(t.Data)
			selfClosing := tt == html.SelfClosingTagToken || voidTags[name]
			parentHidden := hiddenStack[len(hiddenStack)-1]
			thisHidden := parentHidden || tagHidden(t)

			if !selfClosing {
				hiddenStack = append(hiddenStack, thisHidden)
			}

			attrs := attrMap(t)
			hidden := thisHidden || inNoscript

			switch name {
			case "img":
				addImg(attrs["src"], "img-src", hidden)
				addImg(attrs["data-src"], "data-src", hidden)
				addImg(attrs["data-lazy-src"], "data-lazy-src", hidden)
				addImg(attrs["data-original"], "data-original", hidden)
				addImg(attrs["data-lazy"], "data-lazy", hidden)
				addImg(attrs["data-hi-res-src"], "data-hi-res-src", hidden)
				addImg(attrs["data-lowsrc"], "data-lowsrc", hidden)
				addImg(attrs["data-url"], "data-url", hidden)
				addImg(attrs["data-image"], "data-image", hidden)
				addImg(attrs["data-bg"], "data-bg", hidden)
				addImg(attrs["data-background"], "data-background", hidden)
				addImg(attrs["data-background-image"], "data-background-image", hidden)
				addImg(attrs["data-src-retina"], "data-src-retina", hidden)
				addImg(attrs["data-avatar"], "data-avatar", hidden)
				addImg(attrs["data-user-image"], "data-user-image", hidden)
				addImg(attrs["data-profile-image"], "data-profile-image", hidden)
				for _, u := range parseSrcset(attrs["srcset"]) {
					addImg(u, "srcset", hidden)
				}
				for _, u := range parseSrcset(attrs["data-srcset"]) {
					addImg(u, "data-srcset", hidden)
				}
				if style := attrs["style"]; style != "" {
					for _, u := range cssURLs(style) {
						addImg(u, "img-style", hidden)
					}
				}
			case "source":
				addImg(attrs["src"], "source-src", hidden)
				addImg(attrs["data-src"], "source-data-src", hidden)
				for _, u := range parseSrcset(attrs["srcset"]) {
					addImg(u, "source-srcset", hidden)
				}
				for _, u := range parseSrcset(attrs["data-srcset"]) {
					addImg(u, "source-data-srcset", hidden)
				}
			case "image", "use":
				addImg(attrs["href"], "svg-href", hidden)
				addImg(attrs["xlink:href"], "svg-xlink", hidden)
				addImg(attrs["src"], "svg-src", hidden)
			case "video":
				addImg(attrs["src"], "video-src", hidden)
				addImg(attrs["data-src"], "video-data-src", hidden)
				addImg(attrs["data-video-src"], "data-video-src", hidden)
				addImg(attrs["poster"], "video-poster", hidden)
			case "input":
				if strings.EqualFold(attrs["type"], "image") {
					addImg(attrs["src"], "input-image", hidden)
				}
			case "object", "embed":
				addImg(attrs["data"], "object-data", hidden)
				addImg(attrs["src"], "embed-src", hidden)
			case "meta":
				prop := strings.ToLower(attrs["property"] + attrs["name"] + attrs["itemprop"])
				if strings.Contains(prop, "image") || strings.Contains(prop, "thumbnail") ||
					strings.Contains(prop, "og:image") || strings.Contains(prop, "twitter:image") ||
					strings.Contains(prop, "msapplication-tileimage") {
					addImg(attrs["content"], "meta-"+prop, hidden)
				}
			case "link":
				rel := strings.ToLower(attrs["rel"])
				as := strings.ToLower(attrs["as"])
				typ := strings.ToLower(attrs["type"])
				href := attrs["href"]
				switch {
				case strings.Contains(rel, "icon") || rel == "apple-touch-icon" || rel == "mask-icon" || rel == "image_src" || rel == "shortcut icon":
					addImg(href, "link-icon", hidden)
				case as == "image" || as == "video" || strings.HasPrefix(typ, "image/") || strings.HasPrefix(typ, "video/"):
					addImg(href, "link-preload-media", hidden)
				case strings.Contains(rel, "stylesheet") || typ == "text/css":
					out.Styles = append(out.Styles, ResourceHit{URL: resolve(page, href), Via: "stylesheet"})
				case strings.Contains(rel, "manifest"):
					out.Manifests = append(out.Manifests, ResourceHit{URL: resolve(page, href), Via: "manifest"})
				case strings.Contains(rel, "alternate") && (strings.Contains(typ, "xml") || strings.Contains(href, "sitemap")):
					out.Sitemaps = append(out.Sitemaps, ResourceHit{URL: resolve(page, href), Via: "link-alternate"})
				}
			case "a":
				href := attrs["href"]
				if href != "" {
					via := "anchor"
					if indexListing {
						via = "index-listing"
					}
					if looksLikeImage(href) {
						addImg(href, via, hidden)
					} else {
						out.Pages = append(out.Pages, ResourceHit{URL: resolve(page, href), Via: via})
					}
				}
			case "iframe", "frame":
				src := attrs["src"]
				if src != "" && !strings.HasPrefix(strings.ToLower(src), "javascript:") {
					out.Pages = append(out.Pages, ResourceHit{URL: resolve(page, src), Via: "iframe"})
				}
			case "script":
				src := attrs["src"]
				if src != "" {
					out.Scripts = append(out.Scripts, ResourceHit{URL: resolve(page, src), Via: "script-src"})
				}
				scriptType = strings.ToLower(attrs["type"])
				inScript = src == ""
				textBuf.Reset()
			case "style":
				inStyle = true
				textBuf.Reset()
			case "noscript":
				inNoscript = true
			}

			for k, v := range attrs {
				if strings.HasPrefix(k, "data-") && looksLikeImage(v) {
					addImg(v, "attr-"+k, hidden)
				}
				if k == "style" {
					for _, u := range cssURLs(v) {
						addImg(u, "style-attr", hidden)
					}
				}
			}

		case html.EndTagToken:
			t := z.Token()
			name := strings.ToLower(t.Data)
			if !voidTags[name] && len(hiddenStack) > 1 {
				hiddenStack = hiddenStack[:len(hiddenStack)-1]
			}
			switch name {
			case "style":
				if inStyle {
					for _, u := range cssURLs(textBuf.String()) {
						addImg(u, "style-tag", hiddenNow())
					}
				}
				inStyle = false
			case "script":
				if inScript {
					extractFromText(page, textBuf.String(), "script-"+scriptType, hiddenNow(), &out, seenImg)
				}
				inScript = false
			case "noscript":
				inNoscript = false
			}

		case html.TextToken:
			if inStyle || inScript {
				textBuf.Write(z.Text())
			}

		case html.CommentToken:
			c := string(z.Text())
			extractFromText(page, c, "html-comment", true, &out, seenImg)
			// Comments often hide full <img> tags.
			inner := extractHTML(page, []byte("<div>"+c+"</div>"))
			for _, img := range inner.Images {
				img.Hidden = true
				if img.Via == "img-src" {
					img.Via = "comment-img"
				} else {
					img.Via = "comment-" + img.Via
				}
				if img.URL != "" && !seenImg[img.URL] {
					seenImg[img.URL] = true
					out.Images = append(out.Images, img)
				}
			}
		}
	}
}

func extractFromText(page *url.URL, text, via string, hidden bool, out *ExtractResult, seen map[string]bool) {
	unescaped := html.UnescapeString(text)
	for _, re := range []*regexp.Regexp{reQuotedImg, reBareImg} {
		for _, m := range re.FindAllStringSubmatch(unescaped, -1) {
			raw := m[0]
			if len(m) > 1 && m[1] != "" {
				raw = m[1]
			}
			addImage(out, seen, page, raw, via, hidden)
		}
	}
	for _, u := range cssURLs(unescaped) {
		addImage(out, seen, page, u, via+"-css-url", hidden)
	}
	for _, m := range reDataURI.FindAllStringSubmatch(unescaped, -1) {
		if len(m) == 4 {
			addDataURI(out, seen, page, m[0], m[2], via+"-data-uri", hidden)
		}
	}
}

func extractCSS(page *url.URL, body []byte) []ImageHit {
	var out ExtractResult
	seen := map[string]bool{}
	extractFromText(page, string(body), "css", false, &out, seen)
	return out.Images
}

func extractJS(page *url.URL, body []byte) []ImageHit {
	var out ExtractResult
	seen := map[string]bool{}
	extractFromText(page, string(body), "javascript", true, &out, seen)
	return out.Images
}

func extractXML(page *url.URL, body []byte) ExtractResult {
	var out ExtractResult
	seen := map[string]bool{}
	text := html.UnescapeString(string(body))
	for _, m := range reXMLLoc.FindAllStringSubmatch(text, -1) {
		raw := strings.TrimSpace(m[1])
		if looksLikeImage(raw) {
			addImage(&out, seen, page, raw, "sitemap-image", false)
		} else {
			low := strings.ToLower(raw)
			if strings.Contains(low, "sitemap") || strings.HasSuffix(low, ".xml") {
				out.Sitemaps = append(out.Sitemaps, ResourceHit{URL: resolve(page, raw), Via: "sitemap-index"})
			} else {
				out.Pages = append(out.Pages, ResourceHit{URL: resolve(page, raw), Via: "sitemap-loc"})
			}
		}
	}
	for _, m := range reMediaURL.FindAllStringSubmatch(text, -1) {
		if looksLikeImage(m[1]) {
			addImage(&out, seen, page, m[1], "xml-attr", false)
		}
	}
	extractFromText(page, text, "xml", false, &out, seen)
	return out
}

func extractJSON(page *url.URL, body []byte) []ImageHit {
	var out ExtractResult
	seen := map[string]bool{}
	extractFromText(page, string(body), "json", true, &out, seen)
	return out.Images
}

func rawPass(page *url.URL, body []byte, out *ExtractResult, seen map[string]bool) {
	extractFromText(page, string(body), "raw-source", true, out, seen)
}

func addImage(out *ExtractResult, seen map[string]bool, page *url.URL, raw, via string, hidden bool) {
	raw = strings.TrimSpace(html.UnescapeString(raw))
	if raw == "" || raw == "#" {
		return
	}
	if strings.HasPrefix(strings.ToLower(raw), "data:image/") || strings.HasPrefix(strings.ToLower(raw), "data:video/") {
		m := reDataURI.FindStringSubmatch(raw)
		if len(m) == 4 {
			addDataURI(out, seen, page, m[0], m[2], via, hidden)
		}
		return
	}
	if strings.HasPrefix(strings.ToLower(raw), "javascript:") || strings.HasPrefix(strings.ToLower(raw), "mailto:") {
		return
	}
	abs := resolve(page, raw)
	if abs == "" || seen[abs] {
		return
	}
	seen[abs] = true
	out.Images = append(out.Images, ImageHit{URL: abs, Page: page.String(), Via: via, Hidden: hidden})
}

func addDataURI(out *ExtractResult, seen map[string]bool, page *url.URL, raw, mime, via string, hidden bool) {
	key := "data:" + mime + ":" + raw
	if seen[key] {
		return
	}
	seen[key] = true
	ext := mimeToExt(mime)
	payload := raw
	if i := strings.Index(raw, ","); i >= 0 {
		payload = raw[i+1:]
	}
	out.Images = append(out.Images, ImageHit{
		URL:    raw[:min(80, len(raw))] + "…",
		Page:   page.String(),
		Via:    via,
		Hidden: hidden,
		Data: []byte(strings.Map(func(r rune) rune {
			if r == '\n' || r == '\r' || r == ' ' || r == '\t' {
				return -1
			}
			return r
		}, payload)),
		Ext: ext,
	})
}

func tagHidden(t html.Token) bool {
	attrs := attrMap(t)
	if _, ok := attrs["hidden"]; ok {
		return true
	}
	if strings.EqualFold(attrs["aria-hidden"], "true") {
		return true
	}
	if reDisplayNone.MatchString(attrs["style"]) {
		return true
	}
	if reHiddenClass.MatchString(attrs["class"]) {
		return true
	}
	if t.Data == "img" {
		if attrs["width"] == "0" || attrs["height"] == "0" || attrs["width"] == "1" || attrs["height"] == "1" {
			return true
		}
	}
	return false
}

func attrMap(t html.Token) map[string]string {
	m := make(map[string]string, len(t.Attr))
	for _, a := range t.Attr {
		m[strings.ToLower(a.Key)] = strings.TrimSpace(a.Val)
	}
	return m
}

func parseSrcset(v string) []string {
	if v == "" {
		return nil
	}
	var out []string
	for _, part := range strings.Split(v, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		fields := strings.Fields(part)
		if len(fields) > 0 {
			out = append(out, fields[0])
		}
	}
	return out
}

func isDirectoryIndex(page *url.URL, body []byte) bool {
	n := min(len(body), 8192)
	s := strings.ToLower(string(body[:n]))
	if strings.Contains(s, "index of") || strings.Contains(s, "directory listing") {
		return true
	}
	if page != nil && strings.HasSuffix(page.Path, "/") && (strings.Contains(s, "[dir]") || strings.Contains(s, "<pre>")) {
		return true
	}
	return false
}

func sniffIndexListing(page *url.URL, body []byte, out *ExtractResult, seen map[string]bool) {
	text := html.UnescapeString(string(body))
	for _, m := range reIndexFile.FindAllStringSubmatch(text, -1) {
		if len(m) < 2 {
			continue
		}
		addImage(out, seen, page, m[1], "index-listing", false)
	}
}

func cssURLs(v string) []string {
	var out []string
	for _, m := range reCSSURL.FindAllStringSubmatch(v, -1) {
		out = append(out, strings.TrimSpace(m[1]))
	}
	return out
}

func resolve(base *url.URL, ref string) string {
	ref = strings.TrimSpace(ref)
	if ref == "" || base == nil {
		return ""
	}
	u, err := url.Parse(ref)
	if err != nil {
		return ""
	}
	if u.Scheme == "data" || u.Scheme == "javascript" || u.Scheme == "mailto" || u.Scheme == "tel" {
		return ""
	}
	abs := base.ResolveReference(u)
	abs.Fragment = ""
	return abs.String()
}

func isPagePath(u *url.URL) bool {
	ext := strings.ToLower(path.Ext(u.Path))
	return pageExt[ext]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
