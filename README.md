# Subset Sum

Assume given a set of N integers, say A1, A2, ..., An, and another value K. Determine if there is a subset of the given set with sum equal to given K.

* 1 <= N <= 20
* -10^8 <= Ai <= 10^8
* -10^8 <= K <= 10^8

## Examples

```
Input:

  N = 4
  A = { 1, 2, 4, 7 }
  K = 13

Output:

   True
 
   There is a subset { 2, 4, 7 } with sum 13
   13 = 2 + 4 + 7
```

```
Input:

  N = 4
  A = { 1, 2, 4, 7 }
  K = 15

Output:

   False
 
   There is no subset with sum 15
```

## Approaches

- [Recursive](./recursive/README.md)

- [Backtrack](./backtrack/README.md)

- [Dynamic Programming](./dp/README.md)

