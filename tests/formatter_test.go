package tests

import (
	"testing"
	"go-reloaded/agents"
)

func TestFormatter(t *testing.T) {
	formatter := agents.NewFormatter()
	
	tests := []struct {
		input    string
		expected string
	}{
		{"hello,... world", "hello, ... world"},
		{"text,...more", "text, ...more"},
		{"  extra   spaces  ", "extra spaces"},
		{"normal text", "normal text"},
	}
	
	for _, test := range tests {
		result := formatter.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}