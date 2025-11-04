package tests

import (
	"testing"
	"go-reloaded/agents"
)

func TestTransformer(t *testing.T) {
	transformer := agents.NewTransformer()
	
	tests := []struct {
		input    string
		expected string
	}{
		{"1E (hex)", "30"},
		{"10 (bin)", "2"},
		{"hello (up)", "HELLO"},
		{"WORLD (low)", "world"},
		{"bridge (cap)", "Bridge"},
		{"so exciting (up, 2)", "SO EXCITING"},
	}
	
	for _, test := range tests {
		result := transformer.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}