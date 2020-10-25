package recursive

import "fmt"

const DEBUG = true

func RecursiveSubsetSum(A []int, K int) bool {
  if DEBUG {
    fmt.Println("===========================================")
    fmt.Printf("K: %d, A: %v\n", K, A)
    fmt.Println("-------------------------------------------")
  }
  r := recursiveSubsetSum(A, K, 0, 0)
  if DEBUG && r { fmt.Println() }
  return r
}

func recursiveSubsetSum(A []int, K int, i int, sum int) bool {
  size := len(A)
  if sum == K {
    if DEBUG { fmt.Printf("!!! FOUND !!! %d : ", K) }
    return true
  }
  if size <= i {
    return false
  }
  if DEBUG { fmt.Printf("i: %d, sum: %d, A[i]: %d\n", i, sum, A[i]) }
  if recursiveSubsetSum(A, K, i + 1, sum) {
    return true
  }
  if K >= sum + A[i] {
    if recursiveSubsetSum(A, K, i + 1, sum + A[i]) {
      if DEBUG { fmt.Printf("%d ", A[i]) }
      return true
    }
  }
  return false
}

