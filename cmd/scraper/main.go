package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/magnushjensen/apple-docs-changelog/internal/api"
	"github.com/magnushjensen/apple-docs-changelog/internal/crawler"
	"github.com/magnushjensen/apple-docs-changelog/internal/parser"
)

func main() {
	// Flags
	path := flag.String("path", "/documentation/devicemanagement", "Documentation path to fetch")
	output := flag.String("output", "docs", "Output directory for markdown files")
	all := flag.Bool("all", false, "Fetch all pages in the documentation tree")
	delay := flag.Duration("delay", 100*time.Millisecond, "Delay between requests per worker")
	workers := flag.Int("workers", 5, "Number of parallel workers for fetching pages")
	flag.Parse()

	client := api.NewClient()

	if *all {
		crawlAll(client, *path, *output, *delay, *workers)
	} else {
		fetchSingle(client, *path, *output)
	}
}

func crawlAll(client *api.Client, basePath, outputDir string, delay time.Duration, workers int) {
	c := crawler.New(client, outputDir, delay, workers)

	stats, err := c.CrawlAll(basePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nDone! Total: %d, Success: %d, Failed: %d\n",
		stats.Total, stats.Success, stats.Failed)
}

func fetchSingle(client *api.Client, path, outputDir string) {
	fmt.Printf("Fetching %s...\n", path)

	doc, err := client.FetchDocumentation(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	markdown := parser.ToMarkdown(doc)
	outPath := pathToFilePath(path, outputDir)

	// Ensure directory exists
	dir := filepath.Dir(outPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outPath, []byte(markdown), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Written to %s\n", outPath)
}

func pathToFilePath(docPath, outputDir string) string {
	cleaned := strings.TrimPrefix(docPath, "/documentation/")
	parts := strings.Split(cleaned, "/")

	if len(parts) == 1 {
		return filepath.Join(outputDir, parts[0], "index.md")
	}

	return filepath.Join(outputDir, cleaned+".md")
}
