#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
mkdir -p dist
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o dist/SiteImageGrabber.exe .
echo "wrote $(pwd)/dist/SiteImageGrabber.exe"
