# Dynamic Programming

## Approach

If we define DP[i][j] as it denotes if there is a subset of sum j with 1 to i-th elements of A, DP[i][j] can be represented with the prebiously calcuated values as follows:

```
* if j > A[i-1] : DP[i][j] = DP[i-1][j]
* if j <= A[i-1]: DP[i][j] = DP[i-1][j] || 
                             DP[i][j - A[i-1]]
```

Because if there exists a subset of sum j with 1 to i-1th elements of A, both DP[i][j] and DP[i][j + A[i-1]] are certainly true.

If current element A[i-1] has value greater than j, the value cannot be added to the subset because the sum will be greater than j.
Also, if the sum of the current subset is greater than j, no further calculation is needed because it does not satfifies the problem. 

DP[i][0] is always true because the sum of the subset {} is always 0.

## Complexity Analysis

* Time Complexity: O(N*K)
* Space Complexity: O(N*K)
