# TASK-003-GrammarFixer

**Purpose:**  
Correct small grammatical mistakes such as changing “a” to “an” before words starting with vowels or “h”.

## Mission
The GrammarFixer improves language flow by adjusting articles where necessary,  
ensuring smoother and more natural English output.


## Steps (TDD)

1. **Define the expected behavior:**  
   - Replace “a” with “an” when followed by a vowel or “h”.  
   - Keep “a” unchanged in all other cases.

2. **Create natural-language test scenarios:**  
   - `"a apple"` → `"an apple"`  
   - `"a book"` → `"a book"`  
   - `"a honest story"` → `"an honest story"`

3. **Refine rules to handle punctuation and multiple spaces.**


## Acceptance Criteria
- All “a/an” corrections are accurate.
- Works regardless of capitalization (“A”, “An”).
- Punctuation and spacing remain consistent.


## Validation
- Compare expected and actual sentences after processing.
- Test edge cases with capital letters and punctuation.
