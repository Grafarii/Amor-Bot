package main

import (
	"crypto/subtle"
	"os"
	"strings"
)

const (
	defaultAdminPassword = "batata"

	safeMaxDepth = 4
	safeMaxPages = 500
	safeMinDelay = 80

	adminMaxDepth = 50
	adminMaxPages = 50000
	adminMaxConc  = 16
	adminMaxBytes = 200 << 20
)

func configuredAdminPassword() string {
	if p := strings.TrimSpace(os.Getenv("GRABBER_ADMIN_PASSWORD")); p != "" {
		return p
	}
	return defaultAdminPassword
}

func passwordOK(got string) bool {
	want := configuredAdminPassword()
	gb := []byte(got)
	wb := []byte(want)
	if len(gb) != len(wb) {
		return false
	}
	return subtle.ConstantTimeCompare(gb, wb) == 1
}

// applyBounds enforces safe crawl limits unless admin testing mode is unlocked.
func applyBounds(cfg *Config, unlocked bool) string {
	if unlocked {
		if cfg.MaxDepth > adminMaxDepth || cfg.MaxDepth < 0 {
			cfg.MaxDepth = adminMaxDepth
		}
		if cfg.MaxPages > adminMaxPages {
			cfg.MaxPages = adminMaxPages
		}
		if cfg.Concurrency > adminMaxConc {
			cfg.Concurrency = adminMaxConc
		}
		if cfg.MaxBytes > adminMaxBytes {
			cfg.MaxBytes = adminMaxBytes
		}
		return "admin testing mode — off-site, robots, and size limits unlocked"
	}
	cfg.SameHost = true
	cfg.RespectRobots = true
	if cfg.MaxDepth > safeMaxDepth {
		cfg.MaxDepth = safeMaxDepth
	}
	if cfg.MaxPages > safeMaxPages {
		cfg.MaxPages = safeMaxPages
	}
	if cfg.DelayMs < safeMinDelay {
		cfg.DelayMs = safeMinDelay
	}
	return "safe mode — same site, robots.txt, depth 4, 500 pages"
}
