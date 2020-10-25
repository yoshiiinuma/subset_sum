package backtrack

import "fmt"

const DEBUG = true

type Stack []int

func (s *Stack) push(v int) {
  *s = append(*s, v)
}

func (s *Stack) pop() (int, bool) {
  i := len(*s)
  if i == 0 {
    return 0, true
  }
  v := (*s)[i-1]
  *s = (*s)[:i-1]
  return v, false
}

func BacktrackSubsetSum(A []int, K int) bool {
  if DEBUG {
    fmt.Println("===========================================")
    fmt.Printf("K: %d, A: %v\n", K, A)
    fmt.Println("-------------------------------------------")
  }
  subset := Stack{}
  r := backtrackSubsetSum(A, K, 0, 0, subset)
  if DEBUG && r { fmt.Println() }
  return r
}

func backtrackSubsetSum(A []int, K int, i int, sum int, subset Stack) bool {
  if DEBUG { fmt.Printf("i: %d, sum: %d: %v\n", i, sum, subset) }
  N := len(A)
  if sum == K {
    if DEBUG { fmt.Printf("!!! FOUND !!! %d : %v\n", K, subset) }
    return true
  }
  if N <= i || sum > K {
    if DEBUG { fmt.Printf("NOT FOUND: i: %d, sum: %d: %v\n", i, sum, subset) }
    return false
  }
  for j := i; j < N; j++ {
    subset.push(A[j])
    if backtrackSubsetSum(A, K, j + 1, sum + A[j], subset) {
      return true
    }
    subset.pop()
  }
  return false
}

