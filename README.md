# Apple Docs Changelog

Tracks daily changes to Apple's developer documentation by scraping their JSON API and converting it to Markdown.

_This repository is in no way affiliated with Apple, and solely pulls down what anybody can access and read officially without logging in to any Apple related website._

## How it works

A GitHub Action runs daily, fetches all Device Management documentation from Apple, and commits any changes. Use git history to see what changed and when.

## Usage

```bash
# Build
go build -o scraper ./cmd/scraper

# Fetch single page
./scraper -path /documentation/devicemanagement

# Fetch all pages
./scraper -all

# Fetch with more workers
./scraper -all -workers 10
```

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-path` | `/documentation/devicemanagement` | Documentation path to fetch |
| `-output` | `docs` | Output directory |
| `-all` | `false` | Fetch all pages in tree |
| `-workers` | `5` | Parallel workers |
| `-delay` | `100ms` | Delay between requests per worker |

## Browse docs

Documentation is in the `docs/` folder, mirroring Apple's URL structure.
