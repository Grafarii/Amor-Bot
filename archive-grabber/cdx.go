package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultCDX  = "https://web.archive.org/cdx/search/cdx"
	defaultWayb = "https://web.archive.org"
	archiveUA   = "LuminaArchive/1.0 (personal collection of public Wayback snapshots)"
	extFilter   = `original:.*\.(png|jpe?g|jfif|gif|webp|svg|svgz|bmp|ico|avif|tiff?|heic|heif|jxl|mp4|m4v|webm|mov|mkv|ogv)`
)

type CDXRecord struct {
	Timestamp string
	Original  string
	MIME      string
	Status    string
	Digest    string
	Length    string
}

func snapshotURL(wayback, timestamp, original string) string {
	ts := strings.TrimSpace(timestamp)
	orig := strings.TrimSpace(original)
	if orig == "" {
		return ""
	}
	base := strings.TrimRight(wayback, "/")
	if ts == "" {
		return base + "/web/id_/" + orig
	}
	return base + "/web/" + ts + "id_/" + orig
}

func latestSnapshotURL(wayback, original string) string {
	orig := strings.TrimSpace(original)
	if orig == "" {
		return ""
	}
	return strings.TrimRight(wayback, "/") + "/web/id_/" + orig
}

func queryCDX(ctx context.Context, client *http.Client, cdxBase, target, matchType string, extraFilters []string, limit int) ([]CDXRecord, error) {
	if limit < 1 {
		limit = 200
	}
	if limit > 5000 {
		limit = 5000
	}
	if matchType == "" {
		matchType = "prefix"
	}
	q := url.Values{}
	q.Set("url", target)
	q.Set("matchType", matchType)
	q.Set("output", "json")
	q.Set("fl", "timestamp,original,mimetype,statuscode,digest,length")
	q.Set("filter", "statuscode:200")
	for _, f := range extraFilters {
		if f != "" {
			q.Add("filter", f)
		}
	}
	q.Set("collapse", "digest")
	q.Set("limit", fmt.Sprintf("%d", limit))
	endpoint := strings.TrimRight(cdxBase, "/") + "?" + q.Encode()

	var lastErr error
	for attempt := 0; attempt < 2; attempt++ {
		if attempt > 0 {
			wait := time.Duration(attempt) * 2 * time.Second
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(wait):
			}
		}
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", archiveUA)
		req.Header.Set("Accept", "application/json,text/plain;q=0.9")
		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}
		body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
		resp.Body.Close()
		if err != nil {
			lastErr = err
			continue
		}
		if resp.StatusCode == 429 || resp.StatusCode == 445 || resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("CDX HTTP %d", resp.StatusCode)
			continue
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			return nil, fmt.Errorf("CDX HTTP %d", resp.StatusCode)
		}
		return parseCDXJSON(body)
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("CDX request failed")
	}
	return nil, lastErr
}

func parseCDXJSON(body []byte) ([]CDXRecord, error) {
	trim := bytes.TrimSpace(body)
	if len(trim) == 0 || bytes.Equal(trim, []byte("[]")) {
		return nil, nil
	}
	if bytes.HasPrefix(trim, []byte("<")) {
		return nil, fmt.Errorf("CDX returned HTML instead of an index")
	}
	var rows [][]string
	if err := json.Unmarshal(trim, &rows); err != nil {
		return parseCDXText(trim)
	}
	var out []CDXRecord
	for i, row := range rows {
		if i == 0 && len(row) > 0 && strings.EqualFold(row[0], "timestamp") {
			continue
		}
		if len(row) < 2 {
			continue
		}
		rec := CDXRecord{Timestamp: row[0], Original: row[1]}
		if len(row) > 2 {
			rec.MIME = row[2]
		}
		if len(row) > 3 {
			rec.Status = row[3]
		}
		if len(row) > 4 {
			rec.Digest = row[4]
		}
		if len(row) > 5 {
			rec.Length = row[5]
		}
		if rec.Original != "" && rec.Timestamp != "" {
			out = append(out, rec)
		}
	}
	return out, nil
}

func parseCDXText(body []byte) ([]CDXRecord, error) {
	var out []CDXRecord
	for _, line := range strings.Split(string(body), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		f := strings.Fields(line)
		if len(f) < 2 {
			continue
		}
		// urlkey timestamp original mimetype status digest length
		if len(f) >= 7 {
			out = append(out, CDXRecord{Timestamp: f[1], Original: f[2], MIME: f[3], Status: f[4], Digest: f[5], Length: f[6]})
			continue
		}
		if f[0] != "" && f[1] != "" {
			out = append(out, CDXRecord{Timestamp: f[0], Original: f[1]})
		}
	}
	return out, nil
}

func siteQuery(start *url.URL) (target, matchType string) {
	host := strings.TrimPrefix(strings.ToLower(start.Hostname()), "www.")
	if host == "" {
		return start.String(), "prefix"
	}
	path := start.Path
	if strings.HasSuffix(path, "*") {
		path = strings.TrimSuffix(path, "*")
	}
	if path == "" || path == "/" {
		return host + "/", "prefix"
	}
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return host + path, "prefix"
}
