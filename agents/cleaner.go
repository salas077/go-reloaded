package agents

import (
	"strings" // For working with strings
)

// Cleaner agent removes extra spaces and fixes punctuation
// This is like a worker that cleans up messy text
type Cleaner struct{}

// NewCleaner creates a new cleaner
// This is like hiring a new cleaning worker
func NewCleaner() *Cleaner {
	return &Cleaner{}
}

// Process takes messy text and makes it clean
// This is the main function that does all the cleaning
func (c *Cleaner) Process(text string) string {
	// First, fix quotes like ' hello ' to 'hello'
	text = c.fixQuotes(text)
	
	// Then, fix punctuation spacing like "hello , world" to "hello, world"
	text = c.fixPunctuation(text)
	
	// Finally, clean up any extra spaces
	text = c.cleanExtraSpaces(text)
	
	// Return the clean text
	return text
}

// fixQuotes cleans up quotes by removing extra spaces inside them
// Example: " ' hello world ' " becomes " 'hello world' "
func (c *Cleaner) fixQuotes(text string) string {
	// We need to find quotes and clean the spaces inside them
	result := ""
	insideQuote := false
	quoteContent := ""
	
	// Go through each character in the text
	for _, char := range text {
		if char == '\'' { // Found a quote mark
			if !insideQuote {
				// Starting a quote
				insideQuote = true
				quoteContent = ""
				result += "'" // Add the opening quote
			} else {
				// Ending a quote
				insideQuote = false
				// Clean the content inside quotes (remove extra spaces)
				cleanContent := strings.TrimSpace(quoteContent)
				result += cleanContent + "'" // Add clean content and closing quote
			}
		} else {
			if insideQuote {
				// We're inside quotes, collect the content
				quoteContent += string(char)
			} else {
				// We're outside quotes, just add the character
				result += string(char)
			}
		}
	}
	
	return result
}

// fixPunctuation removes spaces before punctuation and adds spaces after
// Example: "hello , world !" becomes "hello, world!"
func (c *Cleaner) fixPunctuation(text string) string {
	// List of punctuation marks we need to fix
	punctuationMarks := []string{".", ",", "!", "?", ":", ";"}
	
	// Fix each punctuation mark
	for _, punct := range punctuationMarks {
		// Remove spaces before punctuation
		// Example: "hello , world" becomes "hello, world"
		text = strings.ReplaceAll(text, " "+punct, punct)
		
		// Add space after punctuation (if there isn't one already)
		// But don't add space if it's at the end of text
		text = c.addSpaceAfterPunctuation(text, punct)
	}
	
	return text
}

// addSpaceAfterPunctuation adds a space after punctuation if needed
func (c *Cleaner) addSpaceAfterPunctuation(text, punct string) string {
	result := ""
	
	// Go through the text and look for our punctuation mark
	for i := 0; i < len(text); i++ {
		char := string(text[i])
		result += char
		
		// If we found our punctuation mark
		if char == punct {
			// Check if there's a next character
			if i+1 < len(text) {
				nextChar := string(text[i+1])
				// If next character is not a space and not punctuation
				if nextChar != " " && !c.isPunctuation(nextChar) {
					result += " " // Add a space
				}
			}
		}
	}
	
	return result
}

// isPunctuation checks if a character is punctuation
func (c *Cleaner) isPunctuation(char string) bool {
	punctuationMarks := []string{".", ",", "!", "?", ":", ";"}
	for _, punct := range punctuationMarks {
		if char == punct {
			return true
		}
	}
	return false
}

// cleanExtraSpaces removes extra spaces (multiple spaces become single space)
// Example: "hello    world" becomes "hello world"
func (c *Cleaner) cleanExtraSpaces(text string) string {
	// Keep replacing double spaces with single spaces until no more double spaces
	for strings.Contains(text, "  ") {
		text = strings.ReplaceAll(text, "  ", " ")
	}
	
	// Remove spaces at the beginning and end
	return strings.TrimSpace(text)
}