package crawler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/magnushjensen/apple-docs-changelog/internal/api"
	"github.com/magnushjensen/apple-docs-changelog/internal/parser"
)

type Crawler struct {
	client    *api.Client
	outputDir string
	delay     time.Duration
	workers   int
}

type Stats struct {
	Total   int64
	Success int64
	Failed  int64
	Skipped int64
}

func New(client *api.Client, outputDir string, delay time.Duration, workers int) *Crawler {
	if workers < 1 {
		workers = 1
	}
	return &Crawler{
		client:    client,
		outputDir: outputDir,
		delay:     delay,
		workers:   workers,
	}
}

// CrawlAll fetches the hierarchy and crawls all documentation pages
func (c *Crawler) CrawlAll(basePath string) (*Stats, error) {
	fmt.Printf("Fetching hierarchy for %s...\n", basePath)

	hierarchy, err := c.client.FetchHierarchy(basePath)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hierarchy: %w", err)
	}

	// Get the "data" interface language (most common for Apple docs)
	nodes, ok := hierarchy.InterfaceLanguages["data"]
	if !ok {
		// Try "swift" as fallback
		nodes, ok = hierarchy.InterfaceLanguages["swift"]
		if !ok {
			return nil, fmt.Errorf("no supported interface language found in hierarchy")
		}
	}

	// Collect all paths from the hierarchy
	var paths []string
	for _, node := range nodes {
		paths = append(paths, collectPaths(node)...)
	}

	fmt.Printf("Found %d pages to fetch using %d workers\n", len(paths), c.workers)

	stats := &Stats{Total: int64(len(paths))}

	// Create work channel and start workers
	pathsChan := make(chan string, len(paths))
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < c.workers; i++ {
		wg.Add(1)
		go c.worker(i, pathsChan, stats, &wg)
	}

	// Send all paths to the channel
	for _, path := range paths {
		pathsChan <- path
	}
	close(pathsChan)

	// Wait for all workers to finish
	wg.Wait()

	return stats, nil
}

// worker processes paths from the channel
func (c *Crawler) worker(id int, paths <-chan string, stats *Stats, wg *sync.WaitGroup) {
	defer wg.Done()

	for path := range paths {
		current := atomic.LoadInt64(&stats.Success) + atomic.LoadInt64(&stats.Failed)
		fmt.Printf("[%d/%d] Worker %d: %s\n", current+1, stats.Total, id, path)

		if err := c.fetchAndWrite(path); err != nil {
			fmt.Printf("  Worker %d Error: %v\n", id, err)
			atomic.AddInt64(&stats.Failed, 1)
		} else {
			atomic.AddInt64(&stats.Success, 1)
		}

		// Rate limiting per worker
		if c.delay > 0 {
			time.Sleep(c.delay)
		}
	}
}

// collectPaths recursively collects all paths from the hierarchy tree
func collectPaths(node api.HierarchyNode) []string {
	var paths []string

	// Only collect nodes that have a path (skip groupMarkers)
	if node.Path != "" {
		paths = append(paths, node.Path)
	}

	// Recurse into children
	for _, child := range node.Children {
		paths = append(paths, collectPaths(child)...)
	}

	return paths
}

// fetchAndWrite fetches a single page and writes it to disk
func (c *Crawler) fetchAndWrite(path string) error {
	doc, err := c.client.FetchDocumentation(path)
	if err != nil {
		return err
	}

	markdown := parser.ToMarkdown(doc)
	outPath := c.pathToFilePath(path)

	// Ensure directory exists
	dir := filepath.Dir(outPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(outPath, []byte(markdown), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// pathToFilePath converts a documentation path to a file path
func (c *Crawler) pathToFilePath(docPath string) string {
	// Remove /documentation/ prefix
	cleaned := strings.TrimPrefix(docPath, "/documentation/")

	// Split into parts
	parts := strings.Split(cleaned, "/")

	if len(parts) == 1 {
		// Root of a technology -> index.md
		return filepath.Join(c.outputDir, parts[0], "index.md")
	}

	// Nested path -> path.md
	return filepath.Join(c.outputDir, cleaned+".md")
}
