package backtrack

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func TestBacktrackSubsetSum(t *testing.T) {
  var A []int
  var K int

  A = []int{ 2, 3, 5, 7 }
  K = 2
  Equal(t, BacktrackSubsetSum(A, K), true)

  A = []int{ 2, 3, 5, 7 }
  K = 7
  Equal(t, BacktrackSubsetSum(A, K), true)

  A = []int{ 2, 3, 5, 7 }
  K = 17
  Equal(t, BacktrackSubsetSum(A, K), true)

  A = []int{ 2, 3, 5, 7 }
  K = 10
  Equal(t, BacktrackSubsetSum(A, K), true)

  A = []int{ 2, 3, 5, 7 }
  K = 12
  Equal(t, BacktrackSubsetSum(A, K), true)

  A = []int{ 2, 3, 5, 7 }
  K = 14
  Equal(t, BacktrackSubsetSum(A, K), true)

  A = []int{ 2, 3, 5, 7 }
  K = 18
  Equal(t, BacktrackSubsetSum(A, K), false)

  A = []int{ 3, 4, 7 }
  K = 6
  Equal(t, BacktrackSubsetSum(A, K), false)

  A = []int{ 3, 4, 7 }
  K = 15
  Equal(t, BacktrackSubsetSum(A, K), false)
}
