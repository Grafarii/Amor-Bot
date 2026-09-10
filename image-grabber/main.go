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
	flag.StringVar(&cfg.StartURL, "url", "", "the site to collect from")
	flag.StringVar(&cfg.OutDir, "out", "", "folder to save the collection into")
	flag.IntVar(&cfg.MaxDepth, "depth", cfg.MaxDepth, "how many link hops to follow on this site")
	flag.IntVar(&cfg.MaxPages, "max-pages", cfg.MaxPages, "maximum HTML pages to visit on this site")
	flag.BoolVar(&cfg.SameHost, "same-host", cfg.SameHost, "stay on this site's pages (CDN images the site uses are still collected)")
	flag.BoolVar(&cfg.ParseSitemap, "sitemap", cfg.ParseSitemap, "read sitemap.xml and robots.txt sitemaps")
	flag.BoolVar(&cfg.ParseCSS, "css", cfg.ParseCSS, "scan stylesheets for url() images")
	flag.BoolVar(&cfg.ParseJS, "js", cfg.ParseJS, "scan scripts for image URLs")
	flag.BoolVar(&cfg.ParseComments, "comments", cfg.ParseComments, "keep images found in HTML comments")
	flag.BoolVar(&cfg.SaveDataURI, "data-uri", cfg.SaveDataURI, "save inline data:image embeds")
	flag.BoolVar(&cfg.RespectRobots, "robots", cfg.RespectRobots, "respect robots.txt")
	flag.IntVar(&cfg.DelayMs, "delay", cfg.DelayMs, "milliseconds to wait between requests")
	flag.StringVar(&cfg.Cookies, "cookies", "", "session Cookie header after login")
	flag.StringVar(&cfg.LoginURL, "login-url", "", "optional login form URL")
	flag.StringVar(&cfg.LoginUser, "login-user", "", "username for login-url")
	flag.StringVar(&cfg.LoginPass, "login-pass", "", "password for login-url")
	adminPass := flag.String("admin-password", "", "admin password to unlock testing limits")
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
	unlocked := passwordOK(*adminPass)
	log("info", applyBounds(&cfg, unlocked))
	stats, err := Run(ctx, cfg, log)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("\nCollected %d images (%d hidden) from %d pages of this site → %s\n",
		stats.Saved.Load(), stats.Hidden.Load(), stats.Pages.Load(), cfg.OutDir)
}
