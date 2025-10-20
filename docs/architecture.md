Go Reloaded — Architecture Analysis

 1. Περιγραφή του Προβλήματος

Το project - Go Reloaded - είναι ένα εργαλείο γραμμένο σε Go, που διαβάζει ένα αρχείο κειμένου (input) και δημιουργεί ένα νέο αρχείο (output) όπου το κείμενο έχει “διορθωθεί” σύμφωνα με συγκεκριμένους κανόνες.  
Οι κανόνες περιλαμβάνουν μετατροπές αριθμών (hex/bin → decimal), αλλαγές στη μορφή των λέξεων (uppercase, lowercase, capitalize), καθαρισμό στίξης, εισαγωγικών, και μικρές γραμματικές διορθώσεις (π.χ. a → an πριν από φωνήεν).

Το πρόγραμμα δεν “καταλαβαίνει” σημασιολογικά το κείμενο — απλώς περνάει από διαδοχικά - στάδια μετασχηματισμού - (pipeline) και εφαρμόζει κάθε κανόνα.


 2. Κανόνες Μετασχηματισμού (Transformation Rules)

| Εντολή | Περιγραφή | Παράδειγμα |

| (hex) | Μετατρέπει τον αριθμό πριν από το `(hex)` από δεκαεξαδικό σε δεκαδικό. | `1E (hex)` → `30` |
| (bin) | Μετατρέπει τον αριθμό πριν από το `(bin)` από δυαδικό σε δεκαδικό. | `10 (bin)` → `2` |
| (up) | Κάνει την προηγούμενη λέξη **κεφαλαία**. | `go (up)` → `GO` |
| (low) | Κάνει την προηγούμενη λέξη **πεζά**. | `LOUD (low)` → `loud` |
| (cap) | Κάνει την προηγούμενη λέξη **με κεφαλαίο πρώτο γράμμα**. | `bridge (cap)` → `Bridge` |
| (up, n) / **(low, n)** / **(cap, n)** | Εφαρμόζει τη μετατροπή στις **n προηγούμενες λέξεις**. | `so exciting (up, 2)` → `SO EXCITING` |
| Punctuation | Μαζεύει περιττά κενά γύρω από κόμματα, τελείες, θαυμαστικά κ.λπ. | `boring ,what ?` → `boring, what?` |
| Punctuation Groups | Κρατά ενωμένες ομάδες όπως `...` ή `!?`. | `thinking ... You` → `thinking... You` |
| Quotation Marks (' ') | Τοποθετεί τα εισαγωγικά **κολλητά** στη λέξη ή φράση τους. | `' awesome '` → `'awesome'` |
| a → an | Αν η επόμενη λέξη ξεκινά με **φωνήεν (a, e, i, o, u)** ή **h**, το “a” γίνεται “an”. | `a optimist` → `an optimist` |
| Negative numbers | Αν έχει πρόσημο “-” μπροστά, το διατηρεί. | `-1E (hex)` → `-30` |

---

 3. Αρχιτεκτονική (Pipeline vs FSM)

 Pipeline Model

Το pipeline χωρίζει τη λειτουργία σε **ανεξάρτητα στάδια**.  
Κάθε στάδιο κάνει μία συγκεκριμένη εργασία πάνω στο κείμενο και περνά το αποτέλεσμα στο επόμενο.

Ροή:

input.txt
↓
[1] Read Text
↓
[2] Clean punctuation & quotes
↓
[3] Grammar correction (a→an)
↓
[4] Transform numbers (hex/bin)
↓
[5] Word transforms (up/low/cap)
↓
[6] Write to output.txt


Πλεονεκτήματα:
- Κάθε στάδιο έχει **μία ευθύνη** (Separation of Concerns).  
- Εύκολο testing (δοκιμάζεις κάθε στάδιο ξεχωριστά).  
- Ευελιξία: μπορείς να προσθέσεις στάδιο χωρίς να αλλάξεις τα υπόλοιπα.  
- Απόλυτα ταιριαστό σε “data flow” προγράμματα όπως αυτό.


FSM (Finite State Machine)

Χρησιμοποιείται όταν το πρόγραμμα αλλάζει “κατάσταση” ανάλογα με γεγονότα (state transitions).  
Π.χ. παιχνίδια, διαδραστικά μενού, διαδικασίες με pause/resume.

Στο Go Reloaded δεν υπάρχουν “states” ή επιλογές — μόνο επεξεργασία δεδομένων.  
Άρα το FSM θα ήταν υπερβολικά σύνθετο για αυτό το είδος προβλήματος.

---

Επιλογή: Pipeline

> Επιλέγω το **Pipeline model**, επειδή το πρόβλημα είναι αλυσίδα επεξεργασιών πάνω σε δεδομένα.  
> Έτσι, κάθε στάδιο λειτουργεί ανεξάρτητα και ο κώδικας είναι καθαρός, ευανάγνωστος και εύκολος στη συντήρηση.

---

 4. Golden Test Set (Success Test Cases)

 4.1. Από τα audit examples

| # | Input | Expected Output |

| 1 | `it (cap) was the best of times, it was the worst of times (up) , ... IT WAS THE (low, 3) winter of despair.` | `It was the best of times, it was the worst of TIMES, ... it was the winter of despair.` |
| 2 | `Simply add 42 (hex) and 10 (bin) and you will see the result is 68.` | `Simply add 66 and 2 and you will see the result is 68.` |
| 3 | `There is no greater agony than bearing a untold story inside you.` | `There is no greater agony than bearing an untold story inside you.` |
| 4 | `Punctuation tests are ... kinda boring ,what do you think ?` | `Punctuation tests are... kinda boring, what do you think?` |


4.2. Tricky / Original Test Cases

| # | Περιγραφή | Input | Expected Output |

| 5 | (up, n) περνά πάνω από στίξη | `This is, frankly, very surprising (up, 2)!` | `This is, frankly, VERY SURPRISING!` |
| 6 | Πολλαπλές ετικέτες κοντά | `We saw A2 (hex), then 1111 (bin) (cap) at the show.` | `We saw 162, then 15 (Cap) at the show.` |
| 7 | Εισαγωγικά με πολλές λέξεις και στίξη | `He said: ' this is, truly, amazing ' !` | `He said: 'this is, truly, amazing'!` |
| 8 | “a” πριν από “h” | `It was a historic event.` | `It was an historic event.` |
| 9 | (low, n) σε ανακατεμένο κείμενο | `PLEASE Keep QUIET (low, 3).` | `please keep quiet.` |
| 10 | Συνδυασμός hex + up | `Add -1E (hex) , then shout now (up, 1)!` | `Add -30, then shout NOW!` |

---

 4.3. Μεγάλη Παράγραφος (Full Flow Example)

Input:

As the guide said: ' welcome to the brooklyn bridge (cap) ' … please add 1E (hex) and 10 (bin) , then say go (up) ! There was a unusual vibe, a honest smile, and it was the best (low, 4) EXPERIENCE EVER (low) … what do you think ?

Expected Output:

As the guide said: 'Welcome to the Brooklyn Bridge'… please add 30 and 2, then say GO! There was an unusual vibe, an honest smile, and it was the best experience ever… what do you think?


---
 5. Ροή Εκτέλεσης (Pipeline Flow)

| Στάδιο | Ενέργεια | Παράδειγμα |

| 1. Read | Διαβάζει το input αρχείο. | `"I am a optimist , 1E (hex)"` |
| 2. Clean | Καθαρίζει στίξη και εισαγωγικά. | `"I am a optimist, 1E (hex)"` |
| 3. Grammar | Ελέγχει “a→an”. | `"I am an optimist, 1E (hex)"` |
| 4. Numbers | Μετατρέπει `(hex)` και `(bin)`. | `"I am an optimist, 30"` |
| 5. Word Transform | Εφαρμόζει `(up|low|cap)` και `(…, n)`. | `"I am AN OPTIMIST, 30"` |
| 6. Write | Γράφει το τελικό αποτέλεσμα. | ✅ `"I am an optimist, 30"` |

---

6. Good Practices and Code Style

Το project θα ακολουθεί τους παρακάτω κανόνες καλής γραφής κώδικα:

- DRY (Don't Repeat Yourself)** → Αποφεύγω επανάληψη κώδικα.
- KISS (Keep It Simple Stupid)** → Κρατάω τον κώδικα όσο πιο απλό γίνεται.
- SOC (Separation of Concerns)** → Κάθε function κάνει μόνο μία δουλειά (π.χ. `cleanPunctuation()`, `transformNumbers()`).
- Consistent Naming → Περιγραφικά ονόματα, σταθερό στυλ (camelCase).
- Error Handling → Έλεγχος όλων των σφαλμάτων (`if err != nil { ... }`).
- Formatting → Χρήση `gofmt` και `goimports` για αυτόματη μορφοποίηση.
- Refactoring → Συχνός καθαρισμός/βελτίωση του κώδικα.
- Independent Stages → Κάθε στάδιο του pipeline (Read, Clean, Transform, Write) να είναι πλήρως ανεξάρτητο.



