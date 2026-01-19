package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	BaseURL = "https://developer.apple.com/tutorials/data"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// FetchDocumentation fetches a documentation page by its path
// path should be like "/documentation/devicemanagement"
func (c *Client) FetchDocumentation(path string) (*DocumentationResponse, error) {
	url := fmt.Sprintf("%s%s.json", BaseURL, path)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for %s", resp.StatusCode, url)
	}

	var doc DocumentationResponse
	if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &doc, nil
}

// DocumentationResponse represents the top-level response from Apple's documentation API
type DocumentationResponse struct {
	Identifier             Identifier           `json:"identifier"`
	Metadata               Metadata             `json:"metadata"`
	Abstract               []InlineContent      `json:"abstract"`
	PrimaryContentSections []ContentSection     `json:"primaryContentSections"`
	TopicSections          []TopicSection       `json:"topicSections"`
	References             map[string]Reference `json:"references"`
	Hierarchy              Hierarchy            `json:"hierarchy"`
	SchemaVersion          SchemaVersion        `json:"schemaVersion"`
	Kind                   string               `json:"kind"`
	Variants               []Variant            `json:"variants"`
}

type Identifier struct {
	URL               string `json:"url"`
	InterfaceLanguage string `json:"interfaceLanguage"`
}

type Metadata struct {
	Title       string     `json:"title"`
	Role        string     `json:"role"`
	RoleHeading string     `json:"roleHeading"`
	SymbolKind  string     `json:"symbolKind"`
	ExternalID  string     `json:"externalID"`
	Modules     []Module   `json:"modules"`
	Platforms   []Platform `json:"platforms"`
}

type Module struct {
	Name string `json:"name"`
}

type Platform struct {
	Name         string `json:"name"`
	IntroducedAt string `json:"introducedAt"`
	Beta         bool   `json:"beta"`
}

type ContentSection struct {
	Kind    string    `json:"kind"`
	Content []Content `json:"content"`
}

type Content struct {
	Type          string          `json:"type"`
	Text          string          `json:"text,omitempty"`
	Anchor        string          `json:"anchor,omitempty"`
	Level         int             `json:"level,omitempty"`
	InlineContent []InlineContent `json:"inlineContent,omitempty"`
	Items         []ListItem      `json:"items,omitempty"`
	Syntax        string          `json:"syntax,omitempty"`
	Code          []string        `json:"code,omitempty"`
}

type InlineContent struct {
	Type       string `json:"type"`
	Text       string `json:"text,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	IsActive   bool   `json:"isActive,omitempty"`
	Code       string `json:"code,omitempty"`
}

type ListItem struct {
	Content []Content `json:"content"`
}

type TopicSection struct {
	Title       string   `json:"title"`
	Anchor      string   `json:"anchor"`
	Identifiers []string `json:"identifiers"`
	Generated   bool     `json:"generated,omitempty"`
}

type Reference struct {
	Identifier     string          `json:"identifier"`
	Title          string          `json:"title"`
	URL            string          `json:"url"`
	Kind           string          `json:"kind"`
	Role           string          `json:"role"`
	Type           string          `json:"type"`
	Abstract       []InlineContent `json:"abstract"`
	Fragments      []Fragment      `json:"fragments,omitempty"`
	NavigatorTitle []Fragment      `json:"navigatorTitle,omitempty"`
}

type Fragment struct {
	Kind string `json:"kind"`
	Text string `json:"text"`
}

type Hierarchy struct {
	Paths [][]string `json:"paths"`
}

type SchemaVersion struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

type Variant struct {
	Paths  []string `json:"paths"`
	Traits []Trait  `json:"traits"`
}

type Trait struct {
	InterfaceLanguage string `json:"interfaceLanguage"`
}

// HierarchyResponse represents the documentation tree structure
type HierarchyResponse struct {
	InterfaceLanguages map[string][]HierarchyNode `json:"interfaceLanguages"`
}

// HierarchyNode represents a node in the documentation tree
type HierarchyNode struct {
	Path       string          `json:"path,omitempty"`
	Title      string          `json:"title"`
	Type       string          `json:"type"`
	Children   []HierarchyNode `json:"children,omitempty"`
	Deprecated bool            `json:"deprecated,omitempty"`
}

// FetchHierarchy fetches the documentation hierarchy/index for a technology
func (c *Client) FetchHierarchy(path string) (*HierarchyResponse, error) {
	if strings.Contains(path, "/documentation") {
		path = strings.TrimPrefix(path, "/documentation")
		path = "/index" + path
	}
	url := fmt.Sprintf("%s%s", BaseURL, path)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d for %s", resp.StatusCode, url)
	}

	var hierarchy HierarchyResponse
	if err := json.NewDecoder(resp.Body).Decode(&hierarchy); err != nil {
		return nil, fmt.Errorf("failed to decode hierarchy response: %w", err)
	}

	return &hierarchy, nil
}
