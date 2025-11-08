# TASK-005: Text Transformer Agent

## Objective
Handle all text transformations including number conversions and case changes.

## Requirements
### Number Conversions
- Hex to decimal: `A2 (hex)` → `162`
- Binary to decimal: `1111 (bin)` → `15`
- Handle negative numbers: `-1E (hex)` → `-30`

### Case Transformations
- Uppercase: `word (up)` → `WORD`
- Lowercase: `WORD (low)` → `word`
- Capitalize: `word (cap)` → `Word`

### Numbered Transformations
- Multi-word: `hello world (up, 2)` → `HELLO WORLD`
- Handle insufficient words gracefully
- Support split tags: `(up,` `2)`

## Implementation
- `agents/transformer.go`
- `NewTransformer()` constructor
- `Process(text string) string` method

## Success Criteria
- All transformation types work correctly
- Handles edge cases and errors
- Preserves punctuation during transformations
- Numbers with (cap) become (Cap)

## Dependencies
- Receives grammar-fixed text from GrammarFixer

## Testing
- Unit tests in `tests/transformer_test.go`
- Test all transformation types
- Edge case validation