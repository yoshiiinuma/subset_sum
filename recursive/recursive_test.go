package recursive

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func TestRecursiveSubsetSum(t *testing.T) {
  var A []int
  var K int

  A = []int{ 1, 2, 4, 7 }
  K = 1
  Equal(t, RecursiveSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 7
  Equal(t, RecursiveSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 13
  Equal(t, RecursiveSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 14
  Equal(t, RecursiveSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 15
  Equal(t, RecursiveSubsetSum(A, K), false)

  A = []int{ 3, 4, 7 }
  K = 6
  Equal(t, RecursiveSubsetSum(A, K), false)
}
