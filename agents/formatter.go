package agents

import (
	"regexp"
	"strings"
)

// Formatter agent finalizes text formatting
type Formatter struct{}

func NewFormatter() *Formatter {
	return &Formatter{}
}

func (f *Formatter) Process(text string) string {
	// Handle ellipsis spacing
	text = strings.ReplaceAll(text, ",... ", ", ... ")
	text = strings.ReplaceAll(text, ",...", ", ...")
	
	// Clean up extra spaces
	re := regexp.MustCompile(`\s+`)
	text = re.ReplaceAllString(text, " ")
	
	return strings.TrimSpace(text)
}