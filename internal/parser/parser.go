package parser

import "strings"

const (
	markdownHeadingSymbol = "#"
	headingSep            = " "
	headingURLSep         = "-"
	newlineSep            = "\n"
)

func isHeading(s string) bool {
	return strings.HasPrefix(s, markdownHeadingSymbol)
}

func extractHeadingURL(s string) string {
	words := strings.Split(s, headingSep)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return strings.Join(words[1:], headingURLSep)
}

// TODO: make this parse heading ID in format {#heading-id}
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
