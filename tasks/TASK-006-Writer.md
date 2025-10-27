# TASK-006-Writer

**Purpose:**  
Write the processed text into the specified output file.


## Mission
The Writer is the final stage of the pipeline.  
It takes the fully formatted text and saves it to the destination file without data loss or encoding errors.


## Steps (TDD)

1. **Define the expected behavior:**  
   - Save final text exactly as produced by the Formatter.  
   - Maintain UTF-8 encoding and line breaks.

2. **Create natural-language test scenarios:**  
   - Input text → Output file contains identical content.  
   - No corruption, truncation, or missing characters.



## Acceptance Criteria
- File saved correctly in UTF-8 encoding.  
- Output content matches the final text exactly.  
- No extra blank lines or formatting artifacts.


## Validation
- Open the result file and compare it with the expected output.  
- Confirm correct file creation and write permissions.
