# Go Reloaded — Agents and Pipeline Breakdown

This document explains how the project is divided into smaller, independent agents.  
Each agent is responsible for a single step in the text processing pipeline.

The overall flow is:

Reader → Tokenizer → Cleaner → GrammarFixer → Transformer → Formatter → Writer  
                          ↘  
                          ErrorHandler


## 1. General Idea

The project follows a **Pipeline Architecture**.  
Each stage receives the text from the previous step, performs its operation, and passes the result forward.

This approach makes testing, debugging, and extending the system much easier.  
If something goes wrong, we can isolate the issue in one specific agent instead of the entire flow.

## 2. Agents Overview

| Agent | What it does | Input | Output |
|--------|---------------|--------|--------|
| **Reader** | Reads the input file and returns the text to be processed. | Input file | Raw text |
| **Tokenizer** | Splits the raw text into tokens (words, punctuation, and symbols) so that later stages can process them more accurately. | Raw text | List of tokens |
| **Cleaner** | Removes extra spaces and fixes punctuation or quotes. | Tokens or text | Clean text |
| **GrammarFixer** | Detects and corrects small grammatical issues, such as changing "a" to "an" before vowels or "h". | Clean text | Grammar-corrected text |
| **Transformer** | Applies word transformations: `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` and their numbered variations. | Grammar-corrected text | Transformed text |
| **Formatter** | Adjusts final punctuation, handles groups like `!?` or `...`, and ensures proper spacing. | Transformed text | Formatted text |
| **Writer** | Saves the final processed text to the output file. | Formatted text | Output file |
| **ErrorHandler** | Monitors the pipeline, logs potential issues, and allows the process to continue without crashing. | Any stage output | Recovery and log message |


## 3. Development Workflow (TDD)

Each agent should be developed following **Test-Driven Development (TDD)**:

1. **Define the test** – Describe in natural language what the agent must do and what output is expected.  
2. **Implement the logic** – Develop only what is necessary to make the test pass.  
3. **Refactor** – Simplify or generalize the solution, keeping it clear and maintainable.  

This ensures that each part of the system is predictable, testable, and stable before integration.


## 4. Error Handling Strategy

The **ErrorHandler** works as a safety layer for the entire pipeline.  
If one of the agents encounters an error (for example, invalid input), the ErrorHandler captures the issue, provides contextual error messages, and ensures clean error reporting.

This design provides consistent error formatting and prevents duplicate error messages throughout the system.



## 5. Related Task Files

| Agent | Task File |
|-------|----------|
| Reader | [TASK-001-Reader.md](../tasks/TASK-001-Reader.md) |
| Tokenizer | [TASK-002-Tokenizer.md](../tasks/TASK-002-Tokenizer.md) |
| Cleaner | [TASK-003-Cleaner.md](../tasks/TASK-003-Cleaner.md) |
| GrammarFixer | [TASK-004-GrammarFixer.md](../tasks/TASK-004-GrammarFixer.md) |
| Transformer | [TASK-005-Transformer.md](../tasks/TASK-005-Transformer.md) |
| Formatter | [TASK-006-Formatter.md](../tasks/TASK-006-Formatter.md) |
| Writer | [TASK-007-Writer.md](../tasks/TASK-007-Writer.md) |
| ErrorHandler | [TASK-008-ErrorHandler.md](../tasks/TASK-008-ErrorHandler.md) |



## 6. Notes

- All agents are **independent** and can be tested in isolation.  
- The final pipeline connects them in a clear, ordered sequence.  
- Each test focuses on functionality, not implementation details.  
- This modular design allows future improvements (for example, adding a Logger or Token Analyzer).


 *This document describes the structure and responsibilities of each agent in the Go Reloaded project. It focuses on the logic and interaction between stages, without including code implementation.*