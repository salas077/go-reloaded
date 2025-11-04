# **Go Reloaded — Architecture Analysis**

## **1. Problem Description**

The project **Go Reloaded** is a tool written in Go that reads an input text file and generates a new output file where the text is "corrected" according to specific transformation rules.
These rules include number conversions (hex/bin → decimal), text formatting changes (uppercase, lowercase, capitalize), punctuation and quotation cleanup, and minor grammatical corrections (e.g., replacing "a" with "an" before vowels).

The program does **not** understand the semantic meaning of the text — it simply processes it through a sequence of **transformation stages** (a pipeline), applying each rule step by step.


## **2. Transformation Rules**

| Rule | Description | Example |
|------|-------------|----------|
| **(hex)**                                 | Converts the number before `(hex)` from hexadecimal to decimal.                                 | `1E (hex)` → `30`                      |
| **(bin)**                                 | Converts the number before `(bin)` from binary to decimal.                                      | `10 (bin)` → `2`                       |
| **(up)**                                  | Makes the previous word **uppercase**.                                                          | `go (up)` → `GO`                       |
| **(low)**                                 | Makes the previous word **lowercase**.                                                          | `LOUD (low)` → `loud`                  |
| **(cap)**                                 | Capitalizes the previous word (first letter uppercase).                                         | `bridge (cap)` → `Bridge`              |
| **(up, n)** / **(low, n)** / **(cap, n)** | Applies the transformation to the **previous n words**.                                         | `so exciting (up, 2)` → `SO EXCITING`  |
| **Punctuation**                           | Removes unnecessary spaces around punctuation marks (commas, periods, exclamation marks, etc.). | `boring ,what ?` → `boring, what?`     |
| **Punctuation Groups**                    | Keeps grouped punctuation marks together, such as `...` or `!?`.                                | `thinking ... You` → `thinking... You` |
| **Quotation Marks (' ')**                 | Keeps quotation marks **attached** to the quoted text.                                          | `' awesome '` → `'awesome'`            |
| **a → an**                                | Changes "a" to "an" if the next word starts with a **vowel (a, e, i, o, u)** or **h**.          | `a optimist` → `an optimist`           |
| **Negative Numbers**                      | Preserves the negative sign before numbers.                                                     | `-1E (hex)` → `-30`                    |



## **3. Architecture (Pipeline vs FSM)**

### **Pipeline Model**

The pipeline design divides the program into **independent stages**.
Each stage performs one specific operation on the text and passes the result to the next stage.

**Flow:**

input.txt
↓
[1] Read Text
↓
[2] Tokenize text
↓
[3] Clean punctuation & quotes
↓
[4] Grammar correction (a→an)
↓
[5] Transform numbers (hex/bin)
↓
[6] Word transforms (up/low/cap)
↓
[7] Final formatting
↓
[8] Write to output.txt

**Advantages:**

* Each stage has **a single responsibility** (Separation of Concerns).
* Easy to test (each stage can be tested independently).
* Flexible — new stages can be added without affecting others.
* Perfectly suited for **data flow** programs like this one.


### **FSM (Finite State Machine)**

FSM is used when a program changes its **state** based on events (state transitions),
for example: games, interactive menus, or processes with pause/resume logic.

In **Go Reloaded**, there are no distinct states or interactions — only data processing.
Thus, an FSM would be unnecessarily complex for this problem.



###  **Choice: Pipeline**

> I choose the **Pipeline model** because this problem involves a chain of data transformations.
> Each stage operates independently, resulting in clean, readable, and maintainable code.


## **4. Golden Test Set (Success Test Cases)**

### **4.1. From the Audit Examples**

| # | Input | Expected Output |
|---|-------|----------------|
| 1 | `it (cap) was the best of times, it was the worst of times (up) , ... IT WAS THE (low, 3) winter of despair.` | `It was the best of times, it was the worst of TIMES, ... it was the winter of despair.` |
| 2 | `Simply add 42 (hex) and 10 (bin) and you will see the result is 68.`                                         | `Simply add 66 and 2 and you will see the result is 68.`                                 |
| 3 | `There is no greater agony than bearing a untold story inside you.`                                           | `There is no greater agony than bearing an untold story inside you.`                     |
| 4 | `Punctuation tests are ... kinda boring ,what do you think ?`                                                 | `Punctuation tests are... kinda boring, what do you think?`                              |


### **4.2. Tricky / Original Test Cases**

| # | Description | Input | Expected Output |
|---|-------------|-------|----------------|
| 5  | (up, n) passes over punctuation                     | `This is, frankly, very surprising (up, 2)!`          | `This is, frankly, VERY SURPRISING!`     |
| 6  | Multiple nearby tags                                | `We saw A2 (hex), then 1111 (bin) (cap) at the show.` | `We saw 162, then 15 (Cap) at the show.` |
| 7  | Quotation marks with multiple words and punctuation | `He said: ' this is, truly, amazing ' !`              | `He said: 'this is, truly, amazing'!`    |
| 8  | "a" before "h"                                      | `It was a historic event.`                            | `It was an historic event.`              |
| 9  | (low, n) in mixed text                              | `PLEASE Keep QUIET (low, 3).`                         | `please keep quiet.`                     |
| 10 | Combination of hex + up                             | `Add -1E (hex) , then shout now (up, 1)!`             | `Add -30, then shout NOW!`               |



### **4.3. Full Paragraph Example (Full Flow)**

**Input:**

As the guide said: ' welcome to the brooklyn bridge (cap) ' … please add 1E (hex) and 10 (bin) , then say go (up) ! There was a unusual vibe, a honest smile, and it was the best (low, 4) EXPERIENCE EVER (low) … what do you think ?


**Expected Output:**

As the guide said: 'Welcome to the Brooklyn Bridge'… please add 30 and 2, then say GO! There was an unusual vibe, an honest smile, and it was the best experience ever… what do you think?


## **5. Execution Flow (Pipeline Stages)**

| Stage | Agent | Description | Example |
|-------|-------|-------------|----------|
| **1. Read** | Reader | Reads the input file. | `"I am a optimist , 1E (hex)"` |
| **2. Tokenize** | Tokenizer | Normalizes spacing and tokenizes text. | `"I am a optimist , 1E (hex)"` |
| **3. Clean** | Cleaner | Cleans punctuation and quotation marks. | `"I am a optimist, 1E (hex)"` |
| **4. Grammar** | GrammarFixer | Checks and applies "a→an". | `"I am an optimist, 1E (hex)"` |
| **5. Transform** | Transformer | Converts `(hex)`, `(bin)`, and case transforms. | `"I am an optimist, 30"` |
| **6. Format** | Formatter | Final formatting and cleanup. | `"I am an optimist, 30"` |
| **7. Write** | Writer | Writes the final result to output. | ✅ `result.txt` |
| **8. Error** | ErrorHandler | Handles errors throughout pipeline. | Error logging & recovery |



## **6. Good Practices and Code Style**

The project will follow the following good coding practices:

* **DRY (Don't Repeat Yourself):** Avoid code duplication.
* **KISS (Keep It Simple, Stupid):** Keep logic as simple as possible.
* **SOC (Separation of Concerns):** Each function handles one task only (e.g. `cleanPunctuation()`, `transformNumbers()`).
* **Consistent Naming:** Use descriptive, consistent variable and function names (camelCase).
* **Error Handling:** Always handle errors (`if err != nil { ... }`).
* **Formatting:** Use `gofmt` and `goimports` for automatic code formatting
* **Refactoring:** Regularly clean and improve code structure.
* **Independent Stages:** Each pipeline stage (Read, Clean, Transform, Write) should be fully independent.