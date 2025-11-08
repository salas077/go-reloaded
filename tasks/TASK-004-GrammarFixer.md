# TASK-004: Grammar Fixer Agent

## Objective
Fix basic grammar issues, specifically article corrections (a/an).

## Requirements
- Change "a" to "an" before vowels (a, e, i, o, u)
- Change "a" to "an" before words starting with "h"
- Preserve case: "A" → "An", "a" → "an"
- Handle edge cases and word boundaries

## Implementation
- `agents/grammar_fixer.go`
- `NewGrammarFixer()` constructor
- `Process(text string) string` method

## Success Criteria
- Correct a/an usage before vowels and h
- Maintains proper capitalization
- Doesn't affect other words or contexts
- Handles punctuation correctly

## Dependencies
- Receives cleaned text from Cleaner agent

## Testing
- Unit tests in `tests/grammar_test.go`
- Test vowel and h-word cases
- Case preservation validation