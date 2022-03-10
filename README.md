
# Limited Correspondence
1. Read data from file;
2. Split data into groups of cases:
* *A case contains **pairs** (two space-separated lowercase alphabetic strings) and a **sequence** (a decoded sequence generated consuming the pairs);*

#### For each case:
1. Verify if case **isSolvable**:  
   1.1 - If **FALSE**, add into case.solution a string indicating that this case is impossible to solve `(Ex: "IMPOSSIBLE")`;  
   1.2 - If **TRUE**, decode a sequence and add result into case.solution `(Ex: "abcdefgh")`;
2. Print formatted solution `(Ex: Case n: abcdefgh)`;