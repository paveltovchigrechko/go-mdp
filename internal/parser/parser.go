package parser

import (
	"log"
	"regexp"
	"strings"
)

const (
	markdownHeadingSymbol = "#"
	headingSep            = " "
	headingURLSep         = "-"
	newlineSep            = "\n"
	idPattern             = "{#.+}$"
	headingIdSymbols      = "{#}"
)

// isHeading accepts a string and return if the string is a Markdown header (starts with `#` symbol).
func isHeading(s string) bool {
	return strings.HasPrefix(s, markdownHeadingSymbol)
}

// isHeadingId accepts a string and return if the string is a heading ID in format `{#custom-id}`.
func isHeadingId(s string) bool {
	result, err := regexp.MatchString(idPattern, s)
	if err != nil {
		log.Println(err)
		return false
	}

	return result
}

// stripHeadingId accepts a string and trims the heading ID symbols in it.
// Intended use: accept `{#custom-id}`, return `custom-id`.
func stripHeadingId(s string) string {
	return strings.Trim(s, headingIdSymbols)
}

// extractHeadingURL accepts a string representing a Markdown heading line and returns
// the line URL. For example, for `# This is my heading` it returns "this-is-my-heading".
// If the heading has an ID in format `# Heading text {#custom-id}`, the function returns
// the stripped ID: `custom-id`
func extractHeadingURL(s string) string {
	words := strings.Split(s, headingSep)
	lastWord := words[len(words)-1]
	if isHeadingId(lastWord) {
		return stripHeadingId(lastWord)
	}

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words[1:], headingURLSep)
}

// ParseHeadings accepts a string (a Markdown file content) and returns a slice of heading IDs
// in format `some-text` here. If a heading has an ID (in format `# Some text {#custom-id}`),
// function returns stripped ID (`custom-id`).
func ParseHeadings(s string) []string {
	lines := strings.Split(s, newlineSep)
	headingURLs := make([]string, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if isHeading(line) {
			headingURL := extractHeadingURL(line)
			headingURLs = append(headingURLs, headingURL)
		}
	}

	return headingURLs
}
