package parser

import (
	"fmt"
	"strings"

	"github.com/magnushjensen/apple-docs-changelog/internal/api"
)

// ToMarkdown converts a DocumentationResponse to Markdown format
func ToMarkdown(doc *api.DocumentationResponse) string {
	var sb strings.Builder

	// Title
	sb.WriteString(fmt.Sprintf("# %s\n\n", doc.Metadata.Title))

	// Abstract
	if len(doc.Abstract) > 0 {
		sb.WriteString(renderInlineContent(doc.Abstract, doc.References))
		sb.WriteString("\n\n")
	}

	// Platforms
	if len(doc.Metadata.Platforms) > 0 {
		sb.WriteString("**Platforms:** ")
		platforms := make([]string, 0, len(doc.Metadata.Platforms))
		for _, p := range doc.Metadata.Platforms {
			platform := fmt.Sprintf("%s %s", p.Name, p.IntroducedAt)
			if p.Beta {
				platform += " (Beta)"
			}
			platforms = append(platforms, platform)
		}
		sb.WriteString(strings.Join(platforms, ", "))
		sb.WriteString("\n\n")
	}

	// Primary content sections
	for _, section := range doc.PrimaryContentSections {
		if section.Kind == "content" {
			sb.WriteString(renderContent(section.Content, doc.References))
		}
	}

	// Topic sections
	if len(doc.TopicSections) > 0 {
		sb.WriteString("## Topics\n\n")
		for _, topic := range doc.TopicSections {
			sb.WriteString(fmt.Sprintf("### %s\n\n", topic.Title))
			for _, id := range topic.Identifiers {
				if ref, ok := doc.References[id]; ok {
					sb.WriteString(fmt.Sprintf("- [%s](%s)", ref.Title, ref.URL))
					if len(ref.Abstract) > 0 {
						abstract := renderInlineContent(ref.Abstract, doc.References)
						if abstract != "" {
							sb.WriteString(fmt.Sprintf(" - %s", abstract))
						}
					}
					sb.WriteString("\n")
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func renderContent(contents []api.Content, refs map[string]api.Reference) string {
	var sb strings.Builder

	for _, content := range contents {
		switch content.Type {
		case "heading":
			level := content.Level
			if level < 1 {
				level = 2
			}
			sb.WriteString(fmt.Sprintf("%s %s\n\n", strings.Repeat("#", level), content.Text))

		case "paragraph":
			sb.WriteString(renderInlineContent(content.InlineContent, refs))
			sb.WriteString("\n\n")

		case "unorderedList":
			for _, item := range content.Items {
				sb.WriteString("- ")
				sb.WriteString(renderContentList(item.Content, refs))
				sb.WriteString("\n")
			}
			sb.WriteString("\n")

		case "orderedList":
			for i, item := range content.Items {
				sb.WriteString(fmt.Sprintf("%d. ", i+1))
				sb.WriteString(renderContentList(item.Content, refs))
				sb.WriteString("\n")
			}
			sb.WriteString("\n")

		case "codeListing":
			sb.WriteString("```")
			if content.Syntax != "" {
				sb.WriteString(content.Syntax)
			}
			sb.WriteString("\n")
			sb.WriteString(strings.Join(content.Code, "\n"))
			sb.WriteString("\n```\n\n")

		case "aside":
			// Render asides as blockquotes
			sb.WriteString("> ")
			sb.WriteString(renderInlineContent(content.InlineContent, refs))
			sb.WriteString("\n\n")
		}
	}

	return sb.String()
}

func renderContentList(contents []api.Content, refs map[string]api.Reference) string {
	var parts []string
	for _, c := range contents {
		if c.Type == "paragraph" {
			parts = append(parts, renderInlineContent(c.InlineContent, refs))
		} else if c.Type == "text" {
			parts = append(parts, c.Text)
		}
	}
	return strings.Join(parts, " ")
}

func renderInlineContent(inlines []api.InlineContent, refs map[string]api.Reference) string {
	var parts []string

	for _, inline := range inlines {
		switch inline.Type {
		case "text":
			parts = append(parts, inline.Text)

		case "reference":
			if ref, ok := refs[inline.Identifier]; ok {
				if inline.IsActive {
					parts = append(parts, fmt.Sprintf("[%s](%s)", ref.Title, ref.URL))
				} else {
					parts = append(parts, ref.Title)
				}
			}

		case "codeVoice":
			parts = append(parts, fmt.Sprintf("`%s`", inline.Code))

		case "emphasis":
			parts = append(parts, fmt.Sprintf("*%s*", inline.Text))

		case "strong":
			parts = append(parts, fmt.Sprintf("**%s**", inline.Text))

		case "newTerm":
			parts = append(parts, fmt.Sprintf("*%s*", inline.Text))
		}
	}

	return strings.Join(parts, "")
}
