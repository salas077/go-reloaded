package main

import (
	"strings" // For working with strings
)

// processTextWithPipeline processes text using our pipeline
// This is a helper function for our tests
func processTextWithPipeline(pipeline *Pipeline, text string) string {
	// Step 1: Tokenize the text (split into manageable pieces)
	// Ask our tokenizer worker to split the text properly
	tokens := pipeline.tokenizer.Process(text)
	text = strings.Join(tokens, " ") // Put the pieces back together

	// Step 2: Clean punctuation and spacing
	// Ask our cleaner worker to fix punctuation and quotes
	text = pipeline.cleaner.Process(text)

	// Step 3: Fix grammar (change "a" to "an" where needed)
	// Ask our grammar fixer worker to correct grammar
	text = pipeline.grammarFixer.Process(text)

	// Step 4: Apply transformations (hex, binary, case changes)
	// Ask our transformer worker to convert numbers and change case
	text = pipeline.transformer.Process(text)

	// Step 5: Final formatting and cleanup
	// Ask our formatter worker to do final touches
	text = pipeline.formatter.Process(text)

	// Return the final processed text
	return text
}