package agents

import (
	"strconv" // For converting strings to numbers
	"strings" // For working with strings
)

// Transformer agent handles hex, bin, and case transformations
// This is like a worker that changes numbers and text case
type Transformer struct{}

// NewTransformer creates a new transformer
// This is like hiring a new worker
func NewTransformer() *Transformer {
	return &Transformer{}
}

// Process takes text and applies all transformations
// This is the main function that does all the work
func (t *Transformer) Process(text string) string {
	// Keep applying transformations until no more changes happen
	maxIterations := 20 // Prevent infinite loops but allow multiple passes
	for iteration := 0; iteration < maxIterations; iteration++ {
		oldText := text
		
		// Apply transformations in order
		text = t.processHexBin(text)
		text = t.processCaseTransformations(text)
		
		// If nothing changed, we're done
		if text == oldText {
			break
		}
	}
	
	return text
}

// processHexBin handles hex and binary number conversions
func (t *Transformer) processHexBin(text string) string {
	words := strings.Split(text, " ")
	
	for i := 1; i < len(words); i++ {
		// Handle hex conversion (with or without punctuation)
		if words[i] == "(hex)" || (strings.HasPrefix(words[i], "(hex)") && len(words[i]) > 5) {
			// Extract punctuation from the tag
			tagPunctuation := ""
			if len(words[i]) > 5 { // More than just "(hex)"
				tagPunctuation = words[i][5:] // Everything after "(hex)"
			}
			
			// Remove any punctuation from the previous word
			prevWord := words[i-1]
			prevPunctuation := ""
			if len(prevWord) > 0 {
				lastChar := prevWord[len(prevWord)-1:]
				if lastChar == "." || lastChar == "," || lastChar == "!" || lastChar == "?" || lastChar == ":" || lastChar == ";" {
					prevPunctuation = lastChar
					prevWord = prevWord[:len(prevWord)-1]
				}
			}
			
			if val, err := strconv.ParseInt(prevWord, 16, 64); err == nil {
				words[i-1] = strconv.Itoa(int(val)) + prevPunctuation + tagPunctuation
				words[i] = ""
			}
		}
		// Handle binary conversion (with or without punctuation)
		if words[i] == "(bin)" || (strings.HasPrefix(words[i], "(bin)") && len(words[i]) > 5) {
			// Extract punctuation from the tag
			tagPunctuation := ""
			if len(words[i]) > 5 { // More than just "(bin)"
				tagPunctuation = words[i][5:] // Everything after "(bin)"
			}
			
			// Remove any punctuation from the previous word
			prevWord := words[i-1]
			prevPunctuation := ""
			if len(prevWord) > 0 {
				lastChar := prevWord[len(prevWord)-1:]
				if lastChar == "." || lastChar == "," || lastChar == "!" || lastChar == "?" || lastChar == ":" || lastChar == ";" {
					prevPunctuation = lastChar
					prevWord = prevWord[:len(prevWord)-1]
				}
			}
			
			if val, err := strconv.ParseInt(prevWord, 2, 64); err == nil {
				words[i-1] = strconv.Itoa(int(val)) + prevPunctuation + tagPunctuation
				words[i] = ""
			}
		}
	}
	
	return t.cleanJoin(words)
}

// processCaseTransformations handles all case transformations
func (t *Transformer) processCaseTransformations(text string) string {
	words := strings.Split(text, " ")
	
	for i := 1; i < len(words); i++ {
		word := words[i]
		
		// Handle simple case transformations (with or without punctuation)
		if t.isSimpleCaseTag(word, "up") {
			punct := t.extractPunctuation(word)
			words[i-1] = t.transformWord(words[i-1], "up") + punct
			words[i] = ""
		} else if t.isSimpleCaseTag(word, "low") {
			punct := t.extractPunctuation(word)
			words[i-1] = t.transformWord(words[i-1], "low") + punct
			words[i] = ""
		} else if t.isSimpleCaseTag(word, "cap") {
			punct := t.extractPunctuation(word)
			if t.isNumber(words[i-1]) {
				words[i] = "(Cap)" + punct
			} else {
				words[i-1] = t.transformWord(words[i-1], "cap") + punct
				words[i] = ""
			}
		} else if t.isNumberedCaseTag(word, i, words) {
			// Handle numbered transformations like (up, 2) or split (up, and 2)
			return t.applyNumberedTransformation(words, i)
		}
	}
	
	return t.cleanJoin(words)
}

// isNumberedCaseTag checks if current word (and possibly next) form a numbered case tag
func (t *Transformer) isNumberedCaseTag(word string, index int, words []string) bool {
	// Check for complete tag like "(up,2)" or "(low,3)"
	if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") && strings.Contains(word, ",") {
		inside := word[1 : len(word)-1]
		parts := strings.Split(inside, ",")
		if len(parts) == 2 {
			caseType := strings.TrimSpace(parts[0])
			return caseType == "up" || caseType == "low" || caseType == "cap"
		}
	}
	
	// Check for split tag like "(up," followed by "2)" or "2)!"
	if strings.HasPrefix(word, "(") && strings.Contains(word, ",") && index+1 < len(words) {
		nextWord := words[index+1]
		if strings.Contains(nextWord, ")") {
			// Extract case type from word like "(up," 
			commaIndex := strings.Index(word, ",")
			caseType := word[1:commaIndex] // Between ( and ,
			return caseType == "up" || caseType == "low" || caseType == "cap"
		}
	}
	
	return false
}

// applyNumberedTransformation applies numbered case transformations
func (t *Transformer) applyNumberedTransformation(words []string, tagIndex int) string {
	var caseType string
	var count int
	var tagEnd int
	
	word := words[tagIndex]
	
	// Parse the tag - handle both "(up,2)" and "(up," + "2)" formats
	if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") && strings.Contains(word, ",") {
		// Complete tag like "(up,2)"
		inside := word[1 : len(word)-1]
		parts := strings.Split(inside, ",")
		if len(parts) == 2 {
			caseType = strings.TrimSpace(parts[0])
			if c, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
				count = c
				tagEnd = tagIndex
			} else {
				return strings.Join(words, " ")
			}
		} else {
			return strings.Join(words, " ")
		}
	} else if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ",") && tagIndex+1 < len(words) {
		// Split tag like "(up," + "2)" or "(up," + "2)!"
		nextWord := words[tagIndex+1]
		if strings.Contains(nextWord, ")") {
			caseType = word[1 : len(word)-1] // Remove ( and ,
			// Extract number and punctuation from nextWord
			closeIndex := strings.Index(nextWord, ")")
			countStr := nextWord[:closeIndex] // Everything before )
			if c, err := strconv.Atoi(countStr); err == nil {
				count = c
				tagEnd = tagIndex + 1
			} else {
				return strings.Join(words, " ")
			}
		} else {
			return strings.Join(words, " ")
		}
	} else {
		return strings.Join(words, " ")
	}
	
	// Validate case type
	if caseType != "up" && caseType != "low" && caseType != "cap" {
		return strings.Join(words, " ")
	}
	
	// Calculate which words to transform
	startIndex := tagIndex - count
	if startIndex < 0 {
		startIndex = 0
	}
	
	// Check if we have enough words
	if startIndex >= tagIndex {
		// Not enough words - handle special case for cap
		if caseType == "cap" {
			words[tagIndex] = "(Cap)"
			if tagEnd > tagIndex {
				words[tagEnd] = ""
			}
		} else {
			// Remove the tag
			words[tagIndex] = ""
			if tagEnd > tagIndex {
				words[tagEnd] = ""
			}
		}
		return t.cleanJoin(words)
	}
	
	// Check if we're trying to capitalize numbers
	if caseType == "cap" {
		for j := startIndex; j < tagIndex; j++ {
			if t.isNumber(words[j]) {
				// Can't capitalize numbers, treat tag as text
				words[tagIndex] = "(Cap)"
				if tagEnd > tagIndex {
					words[tagEnd] = ""
				}
				return t.cleanJoin(words)
			}
		}
	}
	
	// Apply the transformation
	for j := startIndex; j < tagIndex; j++ {
		words[j] = t.transformWord(words[j], caseType)
	}
	
	// Remove the tag(s) and handle punctuation
	words[tagIndex] = ""
	if tagEnd > tagIndex {
		// Extract any punctuation after the closing parenthesis
		nextWord := words[tagEnd]
		if strings.Contains(nextWord, ")") {
			closeIndex := strings.Index(nextWord, ")")
			if closeIndex < len(nextWord)-1 {
				// There's punctuation after the )
				punctuation := nextWord[closeIndex+1:]
				// Add punctuation to the last transformed word
				if tagIndex > 0 {
					words[tagIndex-1] += punctuation
				}
			}
		}
		words[tagEnd] = ""
	}
	
	return t.cleanJoin(words)
}

// transformWord applies a case transformation to a single word
func (t *Transformer) transformWord(word, caseType string) string {
	if len(word) == 0 {
		return word
	}
	
	// Handle punctuation at the end
	punctuation := ""
	cleanWord := word
	
	// Check if word ends with punctuation
	if len(word) > 0 {
		lastChar := word[len(word)-1:]
		if lastChar == "." || lastChar == "," || lastChar == "!" || lastChar == "?" || lastChar == ":" || lastChar == ";" {
			punctuation = lastChar
			cleanWord = word[:len(word)-1]
		}
	}
	
	// Apply transformation
	switch caseType {
	case "up":
		cleanWord = strings.ToUpper(cleanWord)
	case "low":
		cleanWord = strings.ToLower(cleanWord)
	case "cap":
		if len(cleanWord) > 0 {
			cleanWord = strings.ToUpper(string(cleanWord[0])) + strings.ToLower(cleanWord[1:])
		}
	}
	
	return cleanWord + punctuation
}

// isNumber checks if a word is just a number
func (t *Transformer) isNumber(word string) bool {
	// Remove punctuation first
	cleanWord := word
	if len(word) > 0 {
		lastChar := word[len(word)-1:]
		if lastChar == "." || lastChar == "," || lastChar == "!" || lastChar == "?" || lastChar == ":" || lastChar == ";" {
			cleanWord = word[:len(word)-1]
		}
	}
	
	// Try to convert to number
	_, err := strconv.Atoi(cleanWord)
	return err == nil
}

// isSimpleCaseTag checks if a word is a simple case tag (possibly with punctuation)
func (t *Transformer) isSimpleCaseTag(word, caseType string) bool {
	tag := "(" + caseType + ")"
	return word == tag || strings.HasPrefix(word, tag)
}

// extractPunctuation extracts punctuation from a transformation tag
func (t *Transformer) extractPunctuation(word string) string {
	// For tags like "(up)," or "(low)!" extract the punctuation
	if strings.HasPrefix(word, "(") {
		// Find the closing parenthesis
		closeIndex := strings.Index(word, ")")
		if closeIndex != -1 && closeIndex < len(word)-1 {
			// Return everything after the closing parenthesis
			return word[closeIndex+1:]
		}
	}
	return "" // No punctuation
}

// cleanJoin joins words and cleans up extra spaces
func (t *Transformer) cleanJoin(words []string) string {
	result := strings.Join(words, " ")
	// Clean up any double spaces that might have been created
	for strings.Contains(result, "  ") {
		result = strings.ReplaceAll(result, "  ", " ")
	}
	return strings.TrimSpace(result)
}