package main
func longestPalindromeSubseq(s string) int {
    n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n-1; i >= 0; i-- {
		for j := i+1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = Max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}

//
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
