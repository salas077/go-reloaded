## 📘 Documentation Summary

### `docs/analysis.md`
Explains the project’s logic, rules, and design decision.  
Includes a comparison between **Finite State Machine (FSM)** and **Pipeline Model**,  
and justifies why the Pipeline approach is more efficient for this problem.

### `docs/agents.md`
Describes each **agent** (Tokenizer, Cleaner, GrammarFixer, Transformer, Formatter, Writer, ErrorHandler).  
Explains how they interact step by step, following the principles of modular design and TDD.

### `docs/tests.md`
Lists the official **audit test cases** and additional tricky examples (“edge cases”).  
All tests are described in natural language, showing the expected transformation of text.

##  Tasks Overview

Each file under `/tasks/` describes one development task, written using a **TDD workflow**:
- Define the expected behavior in natural language.
- Describe test cases.
- Set clear acceptance and validation criteria.

| Task ID | Agent | Description |
|----------|--------|-------------|
| 001 | Tokenizer | Split text into tokens (words, punctuation, and symbols). |
| 002 | Cleaner | Remove unnecessary spaces and fix punctuation. |
| 003 | GrammarFixer | Adjust grammar (e.g., “a” → “an”). |
| 004 | Transformer | Apply `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` transformations. |
| 005 | Formatter | Ensure correct punctuation and spacing. |
| 006 | Writer | Save the final text to the output file. |
| 007 | ErrorHandler | Handle and log errors gracefully without interrupting the process. |



##  Workflow Summary

1. **Input file is read** → `Reader`  
2. **Text is split into tokens** → `Tokenizer`  
3. **Cleaned and formatted** → `Cleaner`  
4. **Grammar corrections applied** → `GrammarFixer`  
5. **Transformations executed** → `Transformer`  
6. **Final punctuation checked** → `Formatter`  
7. **Output written to file** → `Writer`  
8. **Errors logged (if any)** → `ErrorHandler`


##  Principles Followed

- **TDD (Test-Driven Development):** Tests defined before implementation.  
- **SOC (Separation of Concerns):** Each agent has one responsibility.  
- **KISS (Keep It Simple):** Minimal logic, maximum clarity.  
- **DRY (Don’t Repeat Yourself):** No duplication across stages.  
- **Error Resilience:** The system continues even if one agent fails.  

##  Goal of This Repository

This repository serves as a **blueprint** for building the Go Reloaded project.  
It defines the logic, flow, and quality criteria — not the implementation itself.  

It demonstrates understanding of:
- software architecture,
- modular design,
- test planning,
- and documentation practices expected in professional development environments.