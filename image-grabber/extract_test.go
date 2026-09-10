package main

import (
	"net/url"
	"strings"
	"testing"
)

func TestExtractHiddenAndBetweenTheLines(t *testing.T) {
	page, _ := url.Parse("https://example.com/gallery/")
	html := `
<!doctype html>
<html>
<head>
  <meta property="og:image" content="/og.jpg">
  <link rel="stylesheet" href="/assets/app.css">
  <link rel="icon" href="/favicon.ico">
  <style>body { background: url("/css-bg.webp"); }</style>
  <script>const hidden = "/from-js.png";</script>
</head>
<body>
  <img src="visible.jpg" alt="shown">
  <img src="lazy.jpg" data-src="lazy-hi.webp" srcset="s1.jpg 1x, s2.jpg 2x">
  <img src="tiny.gif" width="1" height="1">
  <div hidden><img src="attr-hidden.png"></div>
  <div class="sr-only"><img src="screen-reader.png"></div>
  <div style="display:none"><img src="display-none.png"></div>
  <!-- classic hide: <img src="commented.png"> and also https://cdn.example.com/between-lines.avif -->
  <noscript><img src="noscript.jpg"></noscript>
  <a href="/photos/">index</a>
  <a href="/photos/shot.heic">direct</a>
  <script type="application/ld+json">{"image":"/jsonld.jpg"}</script>
</body>
</html>`

	res := extractHTML(page, []byte(html))
	got := map[string]ImageHit{}
	for _, img := range res.Images {
		got[img.URL] = img
	}

	need := []string{
		"https://example.com/gallery/visible.jpg",
		"https://example.com/gallery/lazy-hi.webp",
		"https://example.com/gallery/s2.jpg",
		"https://example.com/gallery/tiny.gif",
		"https://example.com/gallery/attr-hidden.png",
		"https://example.com/gallery/screen-reader.png",
		"https://example.com/gallery/display-none.png",
		"https://example.com/gallery/commented.png",
		"https://cdn.example.com/between-lines.avif",
		"https://example.com/gallery/noscript.jpg",
		"https://example.com/photos/shot.heic",
		"https://example.com/css-bg.webp",
		"https://example.com/from-js.png",
		"https://example.com/jsonld.jpg",
		"https://example.com/og.jpg",
		"https://example.com/favicon.ico",
	}
	for _, u := range need {
		if _, ok := got[u]; !ok {
			t.Errorf("missing %s", u)
		}
	}
	if hit, ok := got["https://example.com/gallery/commented.png"]; !ok || !hit.Hidden {
		t.Errorf("commented.png should be marked hidden")
	}
	if hit, ok := got["https://example.com/gallery/display-none.png"]; !ok || !hit.Hidden {
		t.Errorf("display-none.png should be marked hidden")
	}
	if hit, ok := got["https://example.com/gallery/visible.jpg"]; !ok || hit.Hidden {
		t.Errorf("visible.jpg should not be hidden")
	}

	var hasCSS, hasIndex bool
	for _, s := range res.Styles {
		if strings.HasSuffix(s.URL, "/assets/app.css") {
			hasCSS = true
		}
	}
	for _, p := range res.Pages {
		if strings.HasSuffix(p.URL, "/photos/") {
			hasIndex = true
		}
	}
	if !hasCSS {
		t.Error("expected stylesheet enqueue")
	}
	if !hasIndex {
		t.Error("expected directory/index link enqueue")
	}
}

func TestExtractCSSAndSitemap(t *testing.T) {
	page, _ := url.Parse("https://example.com/assets/app.css")
	css := `/* leftover */ .hero { background-image: url("../img/hero.png"); } .x { content: url('https://example.com/abs.gif'); }`
	imgs := extractCSS(page, []byte(css))
	found := map[string]bool{}
	for _, img := range imgs {
		found[img.URL] = true
	}
	if !found["https://example.com/img/hero.png"] || !found["https://example.com/abs.gif"] {
		t.Fatalf("css urls: %+v", imgs)
	}

	sm, _ := url.Parse("https://example.com/sitemap.xml")
	xml := `<?xml version="1.0"?>
	<urlset>
	  <url><loc>https://example.com/page</loc>
	    <image:loc>https://example.com/sitemap-pic.jpg</image:loc>
	  </url>
	</urlset>`
	res := extractXML(sm, []byte(xml))
	ok := false
	for _, img := range res.Images {
		if img.URL == "https://example.com/sitemap-pic.jpg" {
			ok = true
		}
	}
	if !ok {
		t.Fatalf("sitemap image missing: %+v", res.Images)
	}
}
