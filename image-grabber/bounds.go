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
	adminMaxBytes = 1 << 30
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
		if cfg.MaxBytes < adminMaxBytes {
			cfg.MaxBytes = adminMaxBytes
		}
		cfg.DeepScan = true
		if cfg.Concurrency < 10 {
			cfg.Concurrency = 10
		}
		return "studio unlocked — walk more of this site, take every image it publishes, no queue-full drops"
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
	return "this site — pages stay here; every image the site publishes is collected"
}
