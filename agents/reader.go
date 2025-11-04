package agents

import (
	"os" // For reading files from computer
)

// Reader agent reads input file and returns raw text
// This is like a worker that opens files and reads what's inside
type Reader struct{}

// NewReader creates a new reader
// This is like hiring someone to read files for us
func NewReader() *Reader {
	return &Reader{}
}

// Process opens a file and reads all the text inside it
// filename: the name of the file we want to read (like "input.txt")
// Returns: the text inside the file, or an error if something went wrong
func (r *Reader) Process(filename string) (string, error) {
	// Try to read the file from the computer
	// os.ReadFile opens the file and reads everything inside
	fileData, err := os.ReadFile(filename)
	
	// Check if there was an error reading the file
	if err != nil {
		// Something went wrong (maybe file doesn't exist)
		// Return empty string and the error
		return "", err
	}
	
	// Success! Convert the file data to text and return it
	// fileData is in bytes, we need to convert to string
	text := string(fileData)
	return text, nil // nil means no error
}