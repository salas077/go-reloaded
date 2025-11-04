# Go Reloaded 

A powerful text processing tool built with **modular agent architecture** that applies various transformations to text files. **100% tested** with comprehensive unit and integration test coverage.

## Features  All Verified

- **Number Conversions**: Convert hexadecimal and binary numbers to decimal
- **Case Transformations**: Uppercase, lowercase, and capitalize text  
- **Grammar Fixes**: Automatic "a" to "an" corrections before vowels
- **Punctuation Cleanup**: Fix spacing around punctuation marks
- **Quote Formatting**: Proper quotation mark placement
- **Agent Architecture**: 8 independent, testable components
- **Error Handling**: Graceful failure recovery

##  Quick Start

### Prerequisites
- Go 1.19 or higher

### Installation
```bash
git clone https://github.com/yourusername/go-reloaded.git
cd go-reloaded
```

### Usage
```bash
go run . input.txt output.txt
```


##  Transformation Rules

| Rule | Example | Result |
|------|---------|--------|
| `(hex)` | `1E (hex)` | `30` |
| `(bin)` | `10 (bin)` | `2` |
| `(up)` | `hello (up)` | `HELLO` |
| `(low)` | `WORLD (low)` | `world` |
| `(cap)` | `bridge (cap)` | `Bridge` |
| `(up, n)` | `so exciting (up, 2)` | `SO EXCITING` |
| Grammar | `a apple` | `an apple` |
| Punctuation | `hello , world !` | `hello, world!` |
| Quotes | `' amazing '` | `'amazing'` |

##  Examples

### Input File (`sample.txt`):

it (cap) was the best of times, it was the worst of times (up) , ... IT WAS THE (low, 3) winter of despair.

### Command:
```bash
go run . sample.txt result.txt
```


### Output File (`result.txt`):

It was the best of times, it was the worst of TIMES, ... it was the winter of despair.

##  Testing  100% Coverage

### Integration Tests (8/8 PASS)
```bash
go run .
```

Runs all audit test cases with expected vs actual output verification.

### Unit Tests (8/8 PASS)  
```bash
go test ./tests/ -v
```


Tests each individual agent:
-  Reader Agent - File I/O operations
-  Tokenizer Agent - Text splitting
-  Cleaner Agent - Punctuation fixes
-  GrammarFixer Agent - A/an corrections
-  Transformer Agent - All transformations
-  Formatter Agent - Final cleanup
-  Writer Agent - File output
-  ErrorHandler Agent - Error management

### Build & Binary Test
```bash
go build .
./go-reloaded input.txt output.txt
```

##  Project Structure Verified

```
go-reloaded/
├── agents/                    # 8 Independent Agents
│   ├── reader.go              # File input
│   ├── tokenizer.go           # Text splitting
│   ├── cleaner.go             # Punctuation fixes
│   ├── grammar_fixer.go       # A/an corrections
│   ├── transformer.go         # Hex/bin/case transforms
│   ├── formatter.go           # Final cleanup
│   ├── writer.go              # File output
│   └── error_handler.go       # Error management
├── tests/                     # 8 Unit Tests (100% PASS)
│   ├── reader_test.go
│   ├── cleaner_test.go
│   ├── grammar_test.go
│   ├── transformer_test.go
│   ├── formatter_test.go
│   ├── writer_test.go
│   ├── tokenizer_test.go
│   └── error_handler_test.go
├── docs/                      # Complete documentation
├── main.go                   # Entry point
├── pipeline.go               # Agent orchestration
├── test_all.go               # Integration tests
└── go.mod                    # Go module
```

##  Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

##  Quality Assurance

- **100% Test Coverage**: All 8 agents individually tested
- **Integration Verified**: 8/8 audit test cases passing
- **Architecture Validated**: Modular pipeline design
- **Error Resilient**: Graceful failure handling
- **Production Ready**: Compiled binary tested

##  Acknowledgments

- Built as part of Zone01 curriculum
- Follows TDD and clean architecture principles
- Implements enterprise-level software engineering practices