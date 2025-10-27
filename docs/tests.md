# **Go Reloaded — Golden Test Set**

This document contains all the **success test cases** for the project.
Each test includes the **input text** and the **expected output** after processing through the Go Reloaded tool.
No code is included — these are functional test cases written in natural language.


## **1. Audit Example Tests**
                                                                        
| **1** | `it (cap) was the best of times, it was the worst of times (up) , ... IT WAS THE (low, 3) winter of despair.` | `It was the best of times, it was the worst of TIMES, ... it was the winter of despair.` |
| **2** | `Simply add 42 (hex) and 10 (bin) and you will see the result is 68.`                                         | `Simply add 66 and 2 and you will see the result is 68.`                                 |
| **3** | `There is no greater agony than bearing a untold story inside you.`                                           | `There is no greater agony than bearing an untold story inside you.`                     |
| **4** | `Punctuation tests are ... kinda boring ,what do you think ?`                                                 | `Punctuation tests are... kinda boring, what do you think?`                              |


## **2. Original / Tricky Test Cases**

These are additional tests designed to check edge cases, complex combinations, and special conditions.

| #      | Description                                         | Input                                                 | Expected Output                          |

| **5**  | (up, n) passes over punctuation                     | `This is, frankly, very surprising (up, 2)!`          | `This is, frankly, VERY SURPRISING!`     |
| **6**  | Multiple tags close together                        | `We saw A2 (hex), then 1111 (bin) (cap) at the show.` | `We saw 162, then 15 (Cap) at the show.` |
| **7**  | Quotation marks with multiple words and punctuation | `He said: ' this is, truly, amazing ' !`              | `He said: 'this is, truly, amazing'!`    |
| **8**  | “a” before “h”                                      | `It was a historic event.`                            | `It was an historic event.`              |
| **9**  | (low, n) with mixed text                            | `PLEASE Keep QUIET (low, 3).`                         | `please keep quiet.`                     |
| **10** | Combination of hex + up                             | `Add -1E (hex) , then shout now (up, 1)!`             | `Add -30, then shout NOW!`               |


## **3. Full Paragraph Example**

This test demonstrates the full pipeline flow, combining multiple transformation rules.

### **Input:**


As the guide said: ' welcome to the brooklyn bridge (cap) ' … please add 1E (hex) and 10 (bin) , then say go (up) ! There was a unusual vibe, a honest smile, and it was the best (low, 4) EXPERIENCE EVER (low) … what do you think ?


### **Expected Output:**

As the guide said: 'Welcome to the Brooklyn Bridge'… please add 30 and 2, then say GO! There was an unusual vibe, an honest smile, and it was the best experience ever… what do you think?


## **4. Notes**

* All tests are written in **plain text format**.
* Each case assumes that the Go Reloaded tool processes text **sequentially** through all transformation stages:

  1. Clean punctuation and quotes
  2. Apply grammar corrections (`a → an`)
  3. Transform numbers `(hex)` / `(bin)`
  4. Apply text format transformations `(up/low/cap)`
  5. Preserve punctuation groups and spacing
* These cases are meant to be **success tests** (expected to pass if the tool works correctly).

