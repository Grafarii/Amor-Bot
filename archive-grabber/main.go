package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"strings"
)

//go:embed web/*
var webRoot embed.FS

func main() {
	cfg := defaultConfig()
	gui := flag.Bool("gui", false, "open the desktop UI (default if no URL is given)")
	flag.StringVar(&cfg.StartURL, "url", "", "site to search in the Wayback Machine")
	flag.StringVar(&cfg.OutDir, "out", "", "folder to save the collection into")
	flag.IntVar(&cfg.MaxImages, "max", cfg.MaxImages, "maximum archived images to take")
	flag.BoolVar(&cfg.ScanHTML, "html", cfg.ScanHTML, "also read archived HTML pages for buried images")
	flag.IntVar(&cfg.HTMLPages, "pages", cfg.HTMLPages, "archived HTML pages to scan")
	flag.IntVar(&cfg.DelayMs, "delay", cfg.DelayMs, "milliseconds between Wayback requests")
	flag.Parse()
	if cfg.StartURL == "" && flag.NArg() > 0 {
		cfg.StartURL = flag.Arg(0)
	}
	if cfg.OutDir == "" && flag.NArg() > 1 {
		cfg.OutDir = flag.Arg(1)
	}
	useGUI := *gui || cfg.StartURL == ""
	if useGUI {
		sub, err := fs.Sub(webRoot, "web")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := runGUI(sub); err != nil && !strings.Contains(err.Error(), "closed") {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	if cfg.OutDir == "" {
		cfg.OutDir = "archived-images"
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	log := func(kind, msg string) { fmt.Printf("%-5s %s\n", kind, msg) }
	stats, err := Run(ctx, cfg, log)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("\nCollected %d archived images from %d pages → %s\n",
		stats.Found.Load(), stats.Pages.Load(), cfg.OutDir)
}
