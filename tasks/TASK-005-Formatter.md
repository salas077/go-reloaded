# TASK-005-Formatter

**Purpose:**  
Finalize the structure of the text and ensure proper punctuation and capitalization at the end of processing.

## Mission
The Formatter ensures that after all transformations, the text is consistent and grammatically formatted.  
It prepares the final version for output.



## Steps (TDD)

1. **Define the expected behavior:**  
   - Remove leftover spaces between words and punctuation.  
   - Keep punctuation groups like `!?` and `...` together.  
   - Ensure spacing after commas, semicolons, and question marks.

2. **Create natural-language test scenarios:**  
   - `"hello , world !"` → `"hello, world!"`  
   - `"I was thinking ... you were right"` → `"I was thinking... you were right"`


## Acceptance Criteria
- Punctuation and spacing are 100% clean.  
- Grouped punctuation (`!?`, `...`) remains intact.  
- No redundant spaces at line start or end.


## Validation
- Visually inspect test results.
- Run text comparison to confirm no formatting errors.
