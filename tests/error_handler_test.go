package tests

import (
	"errors"
	"testing"
	"go-reloaded/agents"
)

func TestErrorHandler(t *testing.T) {
	errorHandler := agents.NewErrorHandler()
	
	// Test Handle method (should not panic)
	testError := errors.New("test error")
	errorHandler.Handle(testError, "test context")
	
	// Test HandleWithFallback
	fallback := errorHandler.HandleWithFallback(testError, "test context", "fallback value")
	if fallback != "fallback value" {
		t.Errorf("Expected fallback value, got: %q", fallback)
	}
	
	// Test with nil error
	result := errorHandler.HandleWithFallback(nil, "test context", "fallback value")
	if result != "" {
		t.Errorf("Expected empty string for nil error, got: %q", result)
	}
}