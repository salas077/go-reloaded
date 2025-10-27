# TASK-002-Cleaner

**Purpose:**  
Remove unnecessary spaces and fix punctuation or misplaced quotes.


## Mission
The Cleaner agent ensures the text structure is tidy and grammatically consistent.  
It aligns punctuation next to words, keeps groups like `!?` together, and properly formats `'quotes'`.


## Steps (TDD)

1. **Define the expected behavior:**  
   - Punctuation is always next to the previous word and separated from the next.  
   - Quotes (`'`) are correctly attached around words or phrases.

2. **Create natural-language test scenarios:**  
   - Input: `"hello , world !"` → Output: `"hello, world!"`
   - Input: `" ' amazing ' day "` → Output: `"'amazing' day"`

3. **Develop and refine the logic:**  
   - Handle special cases like `!?` or `...` without breaking them apart.


## Acceptance Criteria
- No extra spaces remain before punctuation.  
- Groups of punctuation are preserved (`!?`, `...`).  
- Quotes are properly formatted around their content.


## Validation
- Review sample outputs visually and confirm they match expected grammar and spacing.
- Check integration with GrammarFixer (no word boundaries lost).
