# Lumina

Desktop studio that collects **every image a website publishes**.

Paste one origin. Lumina walks *that site* — its pages, directory indexes, and sitemaps — and takes every still and motion file it actually uses, including files the site serves from a CDN. It does not wander the rest of the web, guess folder names, or break into anyone’s account.

## What it takes from the site

- `<img src>`, `srcset`, `<picture>` / `<source>`
- `<video src>`, `<source type="video/mp4">`, posters, and direct `.mp4` / `.webm` / `.mov` links
- Every common image extension (JPEG, PNG, GIF, WebP, SVG, BMP, ICO, AVIF, TIFF, HEIC, JPEG XL, PSD, RAW, …)
- Lazy-load attributes (`data-src`, `data-original`, `data-bg`, …)
- CSS `url()` backgrounds (inline and linked stylesheets, including CDN CSS)
- URLs buried in JavaScript and JSON
- HTML comments (`<!-- <img src="secret.png"> -->`)
- Hidden markup (`hidden`, `display:none`, `sr-only`, 1×1 pixels)
- `<noscript>` fallbacks, Open Graph / Twitter images, favicons
- `sitemap.xml` and sitemaps listed in `robots.txt`
- Directory indexes (`Index of /images/`) and the files those listings name
- Inline `data:image/...;base64` embeds

The desktop UI **does not save immediately**. Frames appear in the collection; use **Save selected** or **Save all** to write files. A `_manifest.json` records where each file came from.

Only use Lumina on sites you own or have permission to copy.

## How the walk works

Safe mode stays on **this site’s pages**, respects `robots.txt`, and caps depth at 4 / 500 pages. Images and video the site points at — even on another host — still enter the collection.

Unlock a deeper walk of the same site with the studio key **`batata`**:

- In the UI: Studio key → **Unlock**
- CLI: `-admin-password batata`

Unlocked studio also:

- Follows directory indexes and links found on the page you start from
- Accepts **session cookies** (or a login form you already use) so it can fetch the same pages you can after sign-in
- Uses a browser User-Agent and Referer so hotlink/WAF **403** responses are retried, then skipped instead of failing the run
- Never drops jobs with `queue full`

Override the key with `GRABBER_ADMIN_PASSWORD`.

Lumina does not sign in for you, guess `/users/` or `/media/`, or bypass access controls.

## Run the .exe (Windows)

1. Download `SiteImageGrabber.exe` from this folder’s `dist/` directory (or from the **Build Image Grabber** GitHub Action artifact).
2. Double-click it. A browser window opens on Lumina.
3. Paste the site, such as `https://example.com/` or `https://example.com/images/`.
4. Click **Collect**. Review the collection, then save.

### Command line

```bat
SiteImageGrabber.exe -url https://example.com/ -out C:\Users\You\Downloads\site-images
```

## Build it yourself

From this directory, on any machine with Go 1.22+:

```bash
# Windows .exe (works from Linux/macOS too)
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/SiteImageGrabber.exe .

# Local binary
go build -o lumina .
```

```bat
SiteImageGrabber.exe -url https://example.com/gallery/
```
