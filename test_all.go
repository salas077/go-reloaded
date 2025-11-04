package main

import (
	"fmt"
)

func testAll() {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			"Test 1",
			"it (cap) was the best of times, it was the worst of times (up) , ... IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES, ... it was the winter of despair.",
		},
		{
			"Test 2", 
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			"Test 3",
			"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you.",
		},
		{
			"Test 4",
			"Punctuation tests are ... kinda boring ,what do you think ?",
			"Punctuation tests are... kinda boring, what do you think?",
		},
		{
			"Test 5 - (up, n) passes over punctuation",
			"This is, frankly, very surprising (up, 2)!",
			"This is, frankly, VERY SURPRISING!",
		},
		{
			"Test 6 - Multiple tags",
			"We saw A2 (hex), then 1111 (bin) (cap) at the show.",
			"We saw 162, then 15 (Cap) at the show.",
		},
		{
			"Test 7 - Quotation marks",
			"He said: ' this is, truly, amazing ' !",
			"He said: 'this is, truly, amazing'!",
		},
		{
			"Test 8 - a before h",
			"It was a historic event.",
			"It was an historic event.",
		},
	}

	// Create our pipeline to process text
	pipeline := NewPipeline()
	
	for _, test := range tests {
		// Process the text using our pipeline
		result := processTextWithPipeline(pipeline, test.input)
		if result == test.expected {
			fmt.Printf("✓ %s: PASS\n", test.name)
		} else {
			fmt.Printf("✗ %s: FAIL\n", test.name)
			fmt.Printf("  Expected: %q\n", test.expected)
			fmt.Printf("  Got:      %q\n", result)
		}
	}
}