package tests

import (
	"os"
	"testing"
	"go-reloaded/agents"
)

func TestReader(t *testing.T) {
	// Create test file
	testContent := "Hello world"
	testFile := "test_input.txt"
	
	err := os.WriteFile(testFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(testFile)
	
	// Test reader
	reader := agents.NewReader()
	result, err := reader.Process(testFile)
	
	if err != nil {
		t.Errorf("Reader failed: %v", err)
	}
	
	if result != testContent {
		t.Errorf("Expected %q, got %q", testContent, result)
	}
}