# Recursive

## Approach

We can enumurate all possible subsets, and check whether the sum of each subset equals the given sum K. For each element A[i] in the set, we have two options to: 1) include it, or 2) exclude it. So there exist 2^N subsets. We pick up an element, decide to include it or not in the subset, and checks whether the current subset satisfies the problem. Repeat this process recursively until finding a solution or trying every possible cases.

## Complexity Analysis

* Time Complexity: O(2^N)
* Space Complexity: O(N)

