package agents

import (
	"os"
)

// Writer agent saves processed text to output file
type Writer struct{}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Process(text, filename string) error {
	return os.WriteFile(filename, []byte(text), 0644)
}