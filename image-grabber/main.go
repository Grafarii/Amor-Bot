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
	flag.StringVar(&cfg.StartURL, "url", "", "site or directory-index URL to crawl")
	flag.StringVar(&cfg.OutDir, "out", "", "folder to save images into")
	flag.IntVar(&cfg.MaxDepth, "depth", cfg.MaxDepth, "how many link hops to follow")
	flag.IntVar(&cfg.MaxPages, "max-pages", cfg.MaxPages, "maximum HTML pages to visit")
	flag.BoolVar(&cfg.SameHost, "same-host", cfg.SameHost, "stay on the same hostname")
	flag.BoolVar(&cfg.ParseSitemap, "sitemap", cfg.ParseSitemap, "read sitemap.xml and robots.txt sitemaps")
	flag.BoolVar(&cfg.ParseCSS, "css", cfg.ParseCSS, "scan stylesheets for url() images")
	flag.BoolVar(&cfg.ParseJS, "js", cfg.ParseJS, "scan scripts for image URLs")
	flag.BoolVar(&cfg.ParseComments, "comments", cfg.ParseComments, "keep images found in HTML comments")
	flag.BoolVar(&cfg.SaveDataURI, "data-uri", cfg.SaveDataURI, "save inline data:image embeds")
	flag.BoolVar(&cfg.RespectRobots, "robots", cfg.RespectRobots, "respect robots.txt")
	flag.IntVar(&cfg.DelayMs, "delay", cfg.DelayMs, "milliseconds to wait between requests")
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
		cfg.OutDir = "grabbed-images"
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	log := func(kind, msg string) {
		fmt.Printf("%-5s %s\n", kind, msg)
	}
	stats, err := Run(ctx, cfg, log)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("\nSaved %d images (%d hidden) from %d pages → %s\n",
		stats.Saved.Load(), stats.Hidden.Load(), stats.Pages.Load(), cfg.OutDir)
}
