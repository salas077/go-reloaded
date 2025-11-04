package main

import (
	"go-reloaded/agents" // Import our worker agents
	"strings"            // For joining text
)

// Pipeline is like a factory with different workers in a line
// Each worker does one specific job on the text
type Pipeline struct {
	// These are our workers (agents) that will process the text
	reader       *agents.Reader       // Worker 1: Reads files
	tokenizer    *agents.Tokenizer    // Worker 2: Splits text into pieces
	cleaner      *agents.Cleaner      // Worker 3: Cleans punctuation
	grammarFixer *agents.GrammarFixer // Worker 4: Fixes grammar
	transformer  *agents.Transformer  // Worker 5: Transforms text (hex, case, etc.)
	formatter    *agents.Formatter    // Worker 6: Final formatting
	writer       *agents.Writer       // Worker 7: Writes to output file
	errorHandler *agents.ErrorHandler // Worker 8: Handles errors
}

// NewPipeline creates a new text processing factory
// This is like hiring all our workers and setting up the assembly line
func NewPipeline() *Pipeline {
	return &Pipeline{
		// Hire each type of worker
		reader:       agents.NewReader(),       // Hire a file reader
		tokenizer:    agents.NewTokenizer(),    // Hire a text splitter
		cleaner:      agents.NewCleaner(),      // Hire a text cleaner
		grammarFixer: agents.NewGrammarFixer(), // Hire a grammar fixer
		transformer:  agents.NewTransformer(),  // Hire a text transformer
		formatter:    agents.NewFormatter(),    // Hire a formatter
		writer:       agents.NewWriter(),       // Hire a file writer
		errorHandler: agents.NewErrorHandler(), // Hire an error handler
	}
}

// Process takes an input file and creates an output file with transformed text
// This is like running our factory assembly line
func (p *Pipeline) Process(inputFile, outputFile string) error {
	// Step 1: Read the input file
	// Ask our reader worker to open and read the file
	text, err := p.reader.Process(inputFile)
	if err != nil {
		// Something went wrong reading the file
		p.errorHandler.Fatal(err, "Reader")
		return err // Return the error to main function
	}

	// Step 2: Tokenize the text (split into manageable pieces)
	// Ask our tokenizer worker to split the text properly
	tokens := p.tokenizer.Process(text)
	text = strings.Join(tokens, " ") // Put the pieces back together

	// Step 3: Clean punctuation and spacing
	// Ask our cleaner worker to fix punctuation and quotes
	text = p.cleaner.Process(text)

	// Step 4: Fix grammar (change "a" to "an" where needed)
	// Ask our grammar fixer worker to correct grammar
	text = p.grammarFixer.Process(text)

	// Step 5: Apply transformations (hex, binary, case changes)
	// Ask our transformer worker to convert numbers and change case
	text = p.transformer.Process(text)

	// Step 6: Final formatting and cleanup
	// Ask our formatter worker to do final touches
	text = p.formatter.Process(text)

	// Step 7: Write the result to output file
	// Ask our writer worker to save the final text to a file
	err = p.writer.Process(text, outputFile)
	if err != nil {
		// Something went wrong writing the file
		p.errorHandler.Fatal(err, "Writer")
		return err // Return the error to main function
	}

	// Success! All steps completed without errors
	return nil // nil means no error
}