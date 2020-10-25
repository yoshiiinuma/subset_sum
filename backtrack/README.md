# Backtracking

## Approach

We can incrementaly build a subset by selecting an element one by one, and check whether the sum of the current subset equals the given sum K. If it does not, we remove the element most recently added element from the subset, and try another solution. Repeat this process recursively until finding a solution or trying every possible cases.

## Complexity Analysis

* Time Complexity: O(2^N)
* Space Complexity: O(N)


