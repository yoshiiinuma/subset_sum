
package dp

import (
  "testing"
  . "github.com/stretchr/testify/assert"
)

func TestDpSubsetSum(t *testing.T) {
  var A []int
  var K int

  A = []int{ 1, 2, 4, 7 }
  K = 1
  Equal(t, DpSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 7
  Equal(t, DpSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 13
  Equal(t, DpSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 14
  Equal(t, DpSubsetSum(A, K), true)

  A = []int{ 1, 2, 4, 7 }
  K = 15
  Equal(t, DpSubsetSum(A, K), false)

  A = []int{ 3, 4, 7 }
  K = 6
  Equal(t, DpSubsetSum(A, K), false)

  A = []int{ 3, 5, 7 }
  K = 10
  Equal(t, DpSubsetSum(A, K), true)

  A = []int{ 3, 4, 11, 5, 2 }
  K = 9
  Equal(t, DpSubsetSum(A, K), true)
}
