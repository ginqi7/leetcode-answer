package main
func minDistance(word1 string, word2 string) int {
    n := len(word1)
	m := len(word2)
	
	if n == 0 {
		return m
	}
	
	if m == 0 {
		return n
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = i;
	}
	
	for j := 0; j < m+1; j++ {
		dp[0][j] = j;
	}


	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(dp[i-1][j-1] + 1, dp[i][j-1]+1, dp[i-1][j]+1)
			}

		}
	}
	return dp[n][m]

}

//
func Min(a, b, c int) int {
	min := min(a, b)
	if min > c {
		return c
	}
	return min
}

//
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
