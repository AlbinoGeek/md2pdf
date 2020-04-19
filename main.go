package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

	log "github.com/AlbinoGeek/logxi/v1"
	"github.com/AlbinoGeek/md2pdf/markdown"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [markdown] output\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {
	var (
		css   = flag.String("css", "", "PDF User Stylesheet")
		title = flag.String("title", "md2pdf", "PDF Document Title")
		html  = flag.Bool("html", false, "Save HTML instead of PDF")
	)
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	var r io.ReadCloser
	var outPath string
	if len(flag.Args()) > 1 {
		f, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal("failed to open input file", "error", err)
		}
		r = f
		outPath = flag.Arg(1)
	} else {
		r = os.Stdin
		outPath = flag.Arg(0)
	}
	defer func() { r.Close() }()

	text, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal("failed to read input file", "error", err)
	}

	md := markdown.NewMarkdown(text)

	opts := []string{"--print-media-type"}
	if *css != "" {
		var cssPath string
		u, _ := url.Parse(*css)
		if u.Scheme == "http" || u.Scheme == "https" {
			f, e := ioutil.TempFile(os.TempDir(), "md2pdf-css")
			// defer os.Remove(f.Name())
			if e == nil {
				cssPath = f.Name()
			}

			log.Warn("wkthml requires local file, downloading remote stylesheet",
				"provided URL", *css, "local file", cssPath)

			if err := download(u.String(), cssPath); err != nil {
				log.Fatal("failed to download remote stylesheet file")
			}

			cssPath = "file://" + cssPath
		} else {
			abs, _ := filepath.Abs(*css)
			cssPath = "file://" + abs
		}
		opts = append(opts, "--user-style-sheet", filepath.ToSlash(cssPath))
	}

	if *html {
		if err := ioutil.WriteFile(outPath, md.ToHtml(*title, true), 0644); err != nil {
			log.Fatal("failure writing output file", "error", err)
		}
		return
	}

	if err = md.ToPdf(outPath, *title, opts...); err != nil {
		log.Fatal("failure in HtmlToPdf", "error", err)
	}
}
