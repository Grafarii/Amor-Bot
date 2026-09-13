# Lumina Archive — Wayback image collection

Windows desktop app that searches the **public Internet Archive Wayback Machine** for every snapshot of a site, then lists archived images in a local gallery. Nothing is written to disk until you click **Save selected** or **Save all**.

This is **not** a live-site crawler. It does not log in, bypass paywalls, or fetch files the Archive never stored. It only reads the public CDX index and snapshot URLs.

## Download

`archive-grabber/dist/ArchiveImageGrabber.exe` — double-click, paste a site URL, press **Collect**.

The window talks to `127.0.0.1` on your machine. Quit the app to close the collector.

## What it does

1. Asks the public CDX API (`web.archive.org/cdx/search/cdx`) for archived captures of that site.
2. First pass: every snapshot whose MIME type starts with `image/` (plus known media extensions).
3. Optional second pass: archived HTML pages, parsed for `<img>`, CSS `url()`, and similar — then those image URLs are fetched from the Archive.
4. Shows each unique file in the gallery. Save only when you choose.

Rate-limited by default so the public Archive is not hammered.

## Command line

```text
ArchiveImageGrabber.exe -url https://example.com -out archived-images
ArchiveImageGrabber.exe -url https://example.com -html=false
ArchiveImageGrabber.exe -url https://example.com -max 5000 -delay 200
```

`-html` (default true) also walks archived HTML. `-max` caps how many files land in the gallery. `-delay` is milliseconds between Archive requests.

## Build

```text
go test ./...
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/ArchiveImageGrabber.exe .
```
