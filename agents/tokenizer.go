package agents

import (
	"regexp"
	"strings"
)

// Tokenizer agent splits text into manageable tokens
type Tokenizer struct{}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) Process(text string) []string {
	// Split on spaces but preserve transformation tags
	re := regexp.MustCompile(`\s+`)
	tokens := re.Split(text, -1)
	
	// Filter out empty tokens
	var result []string
	for _, token := range tokens {
		if strings.TrimSpace(token) != "" {
			result = append(result, strings.TrimSpace(token))
		}
	}
	
	return result
}