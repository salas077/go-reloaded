package tests

import (
	"os"
	"testing"
	"go-reloaded/agents"
)

func TestWriter(t *testing.T) {
	writer := agents.NewWriter()
	
	testContent := "Test output content"
	testFile := "test_writer_output.txt"
	
	// Test writing
	err := writer.Process(testContent, testFile)
	if err != nil {
		t.Errorf("Writer failed: %v", err)
	}
	
	// Verify content
	result, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("Failed to read written file: %v", err)
	}
	
	if string(result) != testContent {
		t.Errorf("Expected: %q, Got: %q", testContent, string(result))
	}
	
	// Cleanup
	os.Remove(testFile)
}