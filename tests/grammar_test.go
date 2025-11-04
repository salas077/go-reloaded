package tests

import (
	"testing"
	"go-reloaded/agents"
)

func TestGrammarFixer(t *testing.T) {
	grammar := agents.NewGrammarFixer()
	
	tests := []struct {
		input    string
		expected string
	}{
		{"a apple", "an apple"},
		{"a book", "a book"},
		{"a honest story", "an honest story"},
		{"A elephant", "An elephant"},
	}
	
	for _, test := range tests {
		result := grammar.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}