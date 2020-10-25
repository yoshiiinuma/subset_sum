
package dp

import "fmt"

const DEBUG = true

func DpSubsetSum(A []int, K int) bool {
  if DEBUG {
    fmt.Println("=======================================")
    fmt.Println(A, K)
  }
  N := len(A)
  var DP = make([][]bool, N + 1);
  for i := range DP {
    DP[i] = make([]bool, K + 1);
  }
  for i := 0; i <= N; i++ {
    DP[i][0] = true
  }
  r := dpSubsetSum(DP, A, K)
  return r
}

/**
 * DP[i][j]: will be true if there exists a subset of elements
 *           from A[0...i] with sum j.
 *
 * DP[i][j] = if j > A[i]:
 *              DP[i-1][j]
 *            else:
 *              DP[i-1][j] or
 *              DP[i][j - A[i]
 *
 */
func dpSubsetSum(DP [][]bool, A []int, K int) bool {
  N := len(A);
  var a int
  total := 0
  for i := 1; i <= N; i++ {
    a = A[i - 1]
    total += a
    for j := 1; j <= K && j <= total; j++ {
      if j < a {
        DP[i][j] = DP[i-1][j]
      } else {
        DP[i][j] = DP[i][j - a] || DP[i-1][j]
      }
      if (DEBUG) {
        fmt.Printf("DP[%d][%d] = %t\n", i, j, DP[i][j]);
      }
      if j == K && DP[i][K] {
        return true;
      }
    }
  }
  return false;
}


