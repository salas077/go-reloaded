# TASK-001: File Reader Agent

## Objective
Create an agent responsible for reading input files and providing text content to the pipeline.

## Requirements
- Read text files from filesystem
- Handle file not found errors gracefully
- Return clean text content for processing
- Support various text encodings

## Implementation
- `agents/reader.go`
- `NewReader()` constructor
- `Process(filename string) (string, error)` method

## Success Criteria
- Successfully reads valid text files
- Returns appropriate errors for invalid files
- Handles empty files correctly
- No memory leaks or resource issues

## Dependencies
- None (entry point of pipeline)

## Testing
- Unit tests in `tests/reader_test.go`
- Test with valid files, invalid files, empty files
- Error handling validation