# Limited Correspondence

## Classical problem
Emil, a Polish mathematician, sent a simple puzzle by post to his British friend, Alan. Alan sent a reply saying he didn’t have an infinite amount of time he could spend on such non-essential things. Emil modified his puzzle (making it a bit more restricted) and sent it back to Alan. Alan then solved the puzzle.

Here is the original puzzle Emil sent: given a sequence of pairs of strings <img src="https://render.githubusercontent.com/render/math?math=(a_1, b_1),(a_2, b_2),\ldots,(a_ k, b_ k)">, find a non-empty sequence <img src="https://render.githubusercontent.com/render/math?math=s_1,s_2,\ldots,s_m"> such that the following is true:

<p align="center">
   <img src="https://render.githubusercontent.com/render/math?math=a_{s_1}a_{s_2}\ldots a_{s_ m}=b_{s_1}b_{s_2}\ldots b_{s_ m}">
</p>

Where <img src="https://render.githubusercontent.com/render/math?math=a_{s_1}a_{s_2}\dots"> indicates string concatenation. The modified puzzle that Emil sent added the following restriction: for all <img src="https://render.githubusercontent.com/render/math?math=i%20\neq%20j">, <img src="https://render.githubusercontent.com/render/math?math=s_i%20\neq%20s_j">.

You don’t have enough time to solve Emil’s original puzzle. Can you solve the modified version?

### Input

Input consists of up to 5 test cases, ending at end of file. Each case starts with a line containing an integer <img src="https://render.githubusercontent.com/render/math?math=1%20\leq%20k%20\leq%2011">, followed by k lines. Each of the k lines contains two space-separated lowercase alphabetic strings which represent a pair. Each individual string will be non-empty and at most 100 characters long.

### Output

For each case, display the case number followed by the sequence found (if it is possible to form one) or “IMPOSSIBLE” (if it is not possible to solve the problem). If it is possible but there are multiple sequences, you should prefer the shortest one (in terms of the number of characters output). If there are multiple shortest sequences, choose the one that is lexicographically first. Follow the format of the sample output.

### Font
> [Kattis-Correspondence](https://open.kattis.com/problems/correspondence)

___

## Algorith explained

1. Read data from file;
2. Split data into groups of cases:
* *A case contains **pairs** (two space-separated lowercase alphabetic strings) and a **sequence** (a decoded sequence generated consuming the pairs);*

### For each case:

1. Verify if case **isSolvable**:  
   1.1 - If **FALSE**, add into case.solution a string indicating that this case is impossible to solve `(Ex: "IMPOSSIBLE")`;  
   1.2 - If **TRUE**, decode a sequence and add result into case.solution `(Ex: "abcdefgh")`;
2. Print formatted solution `(Ex: Case n: abcdefgh)`;

___

## Examples

### Sample 1 - Happy Patch

#### Given input

``` 
5
are yo
you u
how nhoware
alan arala
dear de
```

1. Verify if case is solvable using func `isSolvable()`, resulting `TRUE`;
2. Call func `Decode()`;
3. Identify a pair with both strings (A, B) having the same prefix `PAIR 5 -> dear de`;
4. Identify next pair that matches when concatenated with previous sequences (<img src="https://render.githubusercontent.com/render/math?math=a_{s_1}=b_{s_1}">) keeps the same prefix `AR` --> `PAIR 4 -> alan arala`, resulting new sequences <img src="https://render.githubusercontent.com/render/math?math=a_{s_1}%20a_{s_2}"> and <img src="https://render.githubusercontent.com/render/math?math=b_{s_1}%20b_{s_2}"> `dearalan dearala`
5. Identify next pair that matches when concatenated with previous sequences (<img src="https://render.githubusercontent.com/render/math?math=a_{s_1}=b_{s_1}">) keeps the same prefix `N` --> `PAIR 3 -> how nhoware`, resulting new sequences `dearalanhow dearalanhoware`;
6. Identify next pair that matches when concatenated with previous sequences (<img src="https://render.githubusercontent.com/render/math?math=a_{s_1}=b_{s_1}">) keeps the same prefix `ARE` --> `PAIR 1 -> are yo`, resulting new sequences `dearalanhoware dearalanhowareyo`;
7. Identify next pair that matches when concatenated with previous sequences (<img src="https://render.githubusercontent.com/render/math?math=a_{s_1}=b_{s_1}">) keeps the same prefix `YO` --> `PAIR 2 -> you u`, resulting new sequences `dearalanhowareyou dearalanhowareyou`;
8. Return func `getSolution(case Case)`;
9. Print result from func `getSolution(case Case)`;

#### Resulted output
```
Case 1: dearalanhowareyou
```
___
   </td>
   </tr>
   </table>
</div>

<div>
   <table>
   <tr>
   <th> Sample Input 2</th>
   <th> Sample Output 2</th>
   </tr>
   <tr>
   <td>

   ```

   8
   i ie
   ing ding
   resp orres
   ond pon
   oyc y
   hello hi
   enj njo
   or c

   ```

   </td>
   <td>

   ```

   Case 2: ienjoycorresponding

   ```

   </td>
   </tr>
   </table>
</div>

<div>
   <table>
   <tr>
   <th> Sample Input 3</th>
   <th> Sample Output 3</th>
   </tr>
   <tr>
   <td>

   ```

   3
   efgh efgh
   d cd
   abc ab
   ```
   </td>
   <td>

   ```

   Case 3: abcd

   ```

   </td>
   </tr>
   </table>
</div>

<div>
   <table>
   <tr>
   <th> Sample Input 4</th>
   <th> Sample Output 4</th>
   </tr>
   <tr>
   <td>

   ```

   3
   a ab
   b bb
   c cc
   ```
   </td>
   <td>

   ```

   Case 4: IMPOSSIBLE

   ```

   </td>
   </tr>
   </table>
</div>
