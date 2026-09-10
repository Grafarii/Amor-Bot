# Site Image Grabber

Windows app that crawls a website (or a directory index) and downloads **every image it can find** — including ones that are not shown on the page.

## What it collects

- `<img src>`, `srcset`, `<picture>` / `<source>`
- Lazy-load attributes (`data-src`, `data-original`, `data-bg`, …)
- CSS `url()` backgrounds (inline and linked stylesheets)
- URLs buried in JavaScript and JSON
- HTML comments (`<!-- <img src="secret.png"> -->`)
- Hidden markup (`hidden`, `display:none`, `sr-only`, 1×1 pixels)
- `<noscript>` fallbacks, Open Graph / Twitter images, favicons
- `sitemap.xml` and sitemaps listed in `robots.txt`
- Directory indexes (`Index of /images/`) and direct links to image files
- Inline `data:image/...;base64` embeds

A `_manifest.json` file is written next to the downloads so you can see **where** each image came from.

The desktop UI **does not save immediately**. Images appear in a gallery at the bottom; use **Save selected** or **Save all** to write files.

Only use this on sites you own or have permission to copy.

## Admin testing password

Safe mode (default) stays on the same site, respects `robots.txt`, and caps depth at 4 / 500 pages.

Unlock those limits with the admin password **`batata`**:

- In the UI: Admin password → **Unlock limits**
- CLI: `-admin-password batata`

You can override the password with the `GRABBER_ADMIN_PASSWORD` environment variable.

## Run the .exe (Windows)

1. Download `SiteImageGrabber.exe` from this folder’s `dist/` directory (or from the **Build Image Grabber** GitHub Action artifact).
2. Double-click it. A browser window opens with the app UI.
3. Paste a URL such as `https://example.com/` or `https://example.com/images/`.
4. Click **Grab images**. Files land in `grabbed-images` next to the exe.

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
go build -o site-image-grabber .
```

```bat
site-image-grabber.exe -url https://example.com/gallery/
```
