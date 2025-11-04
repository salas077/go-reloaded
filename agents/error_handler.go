package agents

import (
	"fmt"
	"log"
)

// ErrorHandler agent manages errors gracefully
type ErrorHandler struct{}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (e *ErrorHandler) Handle(err error, context string) {
	if err != nil {
		log.Printf("Error in %s: %v", context, err)
	}
}

func (e *ErrorHandler) HandleWithFallback(err error, context, fallback string) string {
	if err != nil {
		log.Printf("Error in %s: %v, using fallback", context, err)
		return fallback
	}
	return ""
}

func (e *ErrorHandler) Fatal(err error, context string) {
	if err != nil {
		fmt.Printf("Fatal error in %s: %v\n", context, err)
	}
}