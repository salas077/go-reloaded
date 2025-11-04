package tests

import (
	"testing"
	"go-reloaded/agents"
)

func TestTokenizer(t *testing.T) {
	tokenizer := agents.NewTokenizer()
	
	tests := []struct {
		input    string
		expected []string
	}{
		{"Hello , world !", []string{"Hello", ",", "world", "!"}},
		{"1E (hex) files", []string{"1E", "(hex)", "files"}},
		{"  multiple   spaces  ", []string{"multiple", "spaces"}},
	}
	
	for _, test := range tests {
		result := tokenizer.Process(test.input)
		if len(result) != len(test.expected) {
			t.Errorf("Input: %q, Expected length: %d, Got length: %d", test.input, len(test.expected), len(result))
			continue
		}
		for i, token := range result {
			if token != test.expected[i] {
				t.Errorf("Input: %q, Expected token[%d]: %q, Got: %q", test.input, i, test.expected[i], token)
			}
		}
	}
}