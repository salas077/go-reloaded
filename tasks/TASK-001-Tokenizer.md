# TASK-001-Tokenizer

**Purpose:**  
Prepare the text for processing by splitting it into small, manageable pieces called tokens.


## Mission
The Tokenizer agent separates words, punctuation, and special markers like `(hex)` or `(up)` into individual tokens.  
This allows all other agents to process the input accurately without confusion from spaces or punctuation.


## Steps (TDD)

1. **Define the expected behavior:**  
   - The Tokenizer should identify words, punctuation, and transformation tags separately.
   - Consecutive spaces or punctuation marks must not break the structure.

2. **Create natural-language test scenarios:**  
   - Input: `"Hello , world !"` → Tokens: `["Hello", ",", "world", "!"]`
   - Input: `"1E (hex) files"` → Tokens: `["1E", "(hex)", "files"]`

3. **Develop and refine the logic:**  
   - Split text safely without losing punctuation.
   - Keep tag markers attached to their purpose (e.g., `(hex)` or `(low, 2)`).



## Acceptance Criteria
- Every word and symbol is correctly separated into tokens.
- Transformation markers are preserved.
- No empty or duplicate tokens.


## Validation
- Run test cases that confirm the tokenizer behaves correctly across simple and complex sentences.
- Check that later stages (Cleaner, GrammarFixer) can work directly with these tokens.
