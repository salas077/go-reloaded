package agents

import (
	"strings" // For working with strings
)

// GrammarFixer agent corrects "a" to "an" before vowels and "h"
// This is like a worker that fixes grammar mistakes
type GrammarFixer struct{}

// NewGrammarFixer creates a new grammar fixer
// This is like hiring a new grammar teacher
func NewGrammarFixer() *GrammarFixer {
	return &GrammarFixer{}
}

// Process takes text and fixes "a" to "an" where needed
// Example: "a apple" becomes "an apple"
func (g *GrammarFixer) Process(text string) string {
	// Split text into words so we can check each one
	words := strings.Split(text, " ")
	
	// Go through each word and check if we need to fix "a" to "an"
	for i := 0; i < len(words)-1; i++ {
		// Get current word and next word
		currentWord := words[i]
		nextWord := words[i+1]
		
		// Check if current word is "a" or "A"
		if currentWord == "a" || currentWord == "A" {
			// Check if next word starts with vowel or "h"
			if g.startsWithVowelOrH(nextWord) {
				// Fix the grammar!
				if currentWord == "A" {
					words[i] = "An" // Keep capital A as capital An
				} else {
					words[i] = "an" // Change lowercase a to lowercase an
				}
			}
		}
	}
	
	// Put all words back together
	return strings.Join(words, " ")
}

// startsWithVowelOrH checks if a word starts with a vowel (a,e,i,o,u) or h
// This helps us know when to use "an" instead of "a"
func (g *GrammarFixer) startsWithVowelOrH(word string) bool {
	// If word is empty, return false
	if len(word) == 0 {
		return false
	}
	
	// Get the first letter and make it lowercase for easier checking
	firstLetter := strings.ToLower(string(word[0]))
	
	// List of letters that need "an" instead of "a"
	vowelsAndH := []string{"a", "e", "i", "o", "u", "h"}
	
	// Check if first letter is in our list
	for _, letter := range vowelsAndH {
		if firstLetter == letter {
			return true // Yes, this word needs "an"
		}
	}
	
	return false // No, this word can use "a"
}