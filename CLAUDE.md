# Apple Docs Changelog

A Go CLI tool that scrapes Apple's developer documentation JSON endpoints and converts them to Markdown files.

## Project Overview

**Goal**: Fetch Apple developer documentation from their JSON API, parse the custom format, and output structured Markdown files. A scheduled GitHub Action runs this daily, committing changes and creating PRs with AI-generated changelogs. Git handles all diffing/history.

**Scope**: Device Management documentation (extensible to other frameworks later)

**Language**: Go (for the CLI tool)

## Apple Documentation API

### Endpoints

- **Base URL**: `https://developer.apple.com`
- **Documentation JSON**: `https://developer.apple.com/tutorials/data/documentation/{path}.json`
- **Hierarchy/Index**: Available at technology root level (e.g., `/tutorials/data/documentation/devicemanagement.json`)

### JSON Response Structures

#### 1. Hierarchy Response (`interfaceLanguages`)

Tree structure representing the documentation navigation. See `device-management-top-leve-hierachy.json` for full example.

```json
{
  "interfaceLanguages": {
    "data": [
      {
        "path": "/documentation/devicemanagement",
        "title": "Device Management",
        "type": "module",
        "children": [
          {
            "title": "Configuration Profiles",
            "type": "groupMarker"
          },
          {
            "path": "/documentation/devicemanagement/some-article",
            "title": "Some Article",
            "type": "article",
            "children": [...]
          }
        ]
      }
    ]
  }
}
```

**Node types**: `module`, `groupMarker`, `article`, `collection`, `dictionary`
**Optional fields**: `children`, `deprecated`

#### 2. Page Content Response

Contains actual documentation content. See `sample-apple-developer-doc-page-with-contents.json` for full example.

**Key fields**:
- `identifier`: URL and interface language
- `metadata`: Title, platforms, modules, role, symbolKind
- `abstract`: Short description (array of inline content)
- `primaryContentSections`: Main content (headings, paragraphs, code, etc.)
- `topicSections`: Related topics grouped by section
- `references`: Dictionary of all referenced documents (key = doc identifier)
- `hierarchy`: Breadcrumb path
- `schemaVersion`: API version (currently 0.3.0)

**Content types in `primaryContentSections`**:
- `heading`: `{ type: "heading", level: 2, anchor: "overview", text: "Overview" }`
- `paragraph`: `{ type: "paragraph", inlineContent: [...] }`
- `reference`: `{ type: "reference", identifier: "doc://...", isActive: true }`
- `text`: `{ type: "text", text: "..." }`

**Platform info** (in `metadata.platforms`):
```json
{
  "name": "iOS",
  "introducedAt": "13.0",
  "beta": false
}
```

## Project Structure (Planned)

```
apple-docs-changelog/
├── cmd/
│   └── scraper/
│       └── main.go          # CLI entry point
├── internal/
│   ├── api/
│   │   └── client.go        # Apple API client
│   └── parser/
│       └── parser.go        # JSON to Markdown conversion
├── docs/                     # Generated Markdown (mirrors Apple hierarchy)
│   └── devicemanagement/
│       ├── index.md
│       ├── configuring-multiple-devices-using-profiles.md
│       ├── profile-specific-payload-keys/
│       │   ├── index.md
│       │   ├── toplevel.md
│       │   └── ...
│       └── ...
├── .github/
│   └── workflows/
│       └── daily-scrape.yml # Scheduled Action
├── go.mod
├── go.sum
└── CLAUDE.md
```

## CLI Commands (Planned)

```bash
# Fetch and convert Device Management docs
./scraper fetch

# Future: specify different documentation sets
./scraper fetch --target devicemanagement
```

## Markdown Output Format

Target structure for converted docs:

```markdown
# {Title}

> {Abstract}

**Platforms**: iOS 13.0+, macOS 10.15+, ...

## Overview

{Content from primaryContentSections}

## Topics

### {Section Title}
- [{Topic Title}]({path}) - {abstract}
```

## Development Notes

- Rate limit API requests appropriately
- Output Markdown files to `docs/` mirroring Apple's URL hierarchy
- Git handles all versioning/diffing - just overwrite files each run
- GitHub Action creates PR with changes; AI generates changelog from diff

## Sample Files

- `sample-apple-developer-doc-page-with-contents.json` - Example page content response
- `device-management-top-leve-hierachy.json` - Example hierarchy/tree response
