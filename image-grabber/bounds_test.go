package main

import (
	"os"
	"testing"
)

func TestApplyBoundsLocksWithoutAdmin(t *testing.T) {
	cfg := defaultConfig()
	cfg.SameHost = false
	cfg.RespectRobots = false
	cfg.MaxDepth = 40
	cfg.MaxPages = 9000
	cfg.DelayMs = 0
	msg := applyBounds(&cfg, false)
	if cfg.SameHost != true || cfg.RespectRobots != true {
		t.Fatalf("safe mode should force same-host and robots: %+v", cfg)
	}
	if cfg.MaxDepth != safeMaxDepth || cfg.MaxPages != safeMaxPages || cfg.DelayMs != safeMinDelay {
		t.Fatalf("safe caps not applied: %+v", cfg)
	}
	if msg == "" {
		t.Fatal("expected status message")
	}
}

func TestApplyBoundsUnlocksWithAdmin(t *testing.T) {
	cfg := defaultConfig()
	cfg.SameHost = false
	cfg.RespectRobots = false
	cfg.MaxDepth = 12
	cfg.MaxPages = 2000
	cfg.DelayMs = 0
	applyBounds(&cfg, true)
	if cfg.SameHost || cfg.RespectRobots || cfg.DelayMs != 0 {
		t.Fatalf("admin mode should keep testing choices: %+v", cfg)
	}
	if cfg.MaxDepth != 12 || cfg.MaxPages != 2000 {
		t.Fatalf("admin mode should keep requested limits: %+v", cfg)
	}
	if !cfg.DeepScan {
		t.Fatal("admin mode should enable deep folder/user scan")
	}
}

func TestPasswordOK(t *testing.T) {
	if defaultAdminPassword != "batata" {
		t.Fatalf("default password want batata, got %q", defaultAdminPassword)
	}
	if passwordOK("nope") || passwordOK("grafarii") {
		t.Fatal("wrong password accepted")
	}
	if !passwordOK("batata") {
		t.Fatal("batata rejected")
	}
	t.Setenv("GRABBER_ADMIN_PASSWORD", "test-unlock")
	if !passwordOK("test-unlock") {
		t.Fatal("env password rejected")
	}
	if passwordOK("batata") {
		t.Fatal("default still accepted after env override")
	}
	os.Unsetenv("GRABBER_ADMIN_PASSWORD")
}
