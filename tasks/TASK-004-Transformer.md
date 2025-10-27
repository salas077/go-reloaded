# TASK-004-Transformer

**Purpose:**  
Handle all transformation commands like `(hex)`, `(bin)`, `(up)`, `(low)`, and `(cap)`.


## Mission
The Transformer agent interprets transformation markers in the text and modifies the related words accordingly.  
It handles both single-word and multi-word transformations (e.g., `(up, 3)`).


## Steps (TDD)

1. **Define the expected behavior:**  
   - `(hex)` → Convert previous word from hexadecimal to decimal.  
   - `(bin)` → Convert from binary to decimal.  
   - `(up)` / `(low)` / `(cap)` → Change text case of previous word(s).

2. **Create natural-language test scenarios:**  
   - `"1E (hex)"` → `"30"`  
   - `"10 (bin)"` → `"2"`  
   - `"ready, set, go (up)"` → `"ready, set, GO"`
   - `"so exciting (up, 2)"` → `"SO EXCITING"`

3. **Include rules for negative numbers and multiple transformations.**



## Acceptance Criteria
- Each transformation tag produces correct and consistent results.  
- Supports negative numbers like `-1E (hex)` → `-30`.  
- Multi-word transformations work as expected.  
- Transformation tags are removed from the final output.


## Validation
- Compare outputs across all audit examples.  
- Confirm that text remains readable and punctuation untouched.
