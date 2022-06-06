package main
func longestCommonSubsequence(text1 string, text2 string) int {
    m := len(text1)
	n := len(text2)
	
	if m == 0 || n == 0 {
		return 0
	}
	
	dp := make([][]int, m+1)
	for idx, _:= range dp {
		dp[idx] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}
	
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}


	for i := 1; i <= m; i++  {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1]  + 1
			} else {
				dp[i][j] = Max(dp[i-1][j-1], dp[i][j-1], dp[i-1][j])
			}
		}

	}
	return dp[m][n]

}

//
func Max(a, b, c int) int {
	ans := a
	if ans < b {
		ans = b
	}
	if ans < c {
		ans = c
	}
	return ans
}
