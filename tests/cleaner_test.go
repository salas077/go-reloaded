package tests

import (
	"testing"
	"go-reloaded/agents"
)

func TestCleaner(t *testing.T) {
	cleaner := agents.NewCleaner()
	
	tests := []struct {
		input    string
		expected string
	}{
		{"hello , world !", "hello, world!"},
		{" ' amazing ' day ", "'amazing'day "},
		{"test ... case", "test... case"},
	}
	
	for _, test := range tests {
		result := cleaner.Process(test.input)
		if result != test.expected {
			t.Errorf("Input: %q, Expected: %q, Got: %q", test.input, test.expected, result)
		}
	}
}