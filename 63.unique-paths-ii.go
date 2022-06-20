package main
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	dp := make([][]int, m)
	initVal := 1
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 {
			initVal = 0
		} 
		dp[i][0] = initVal
	}
	initVal = 1
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			initVal = 0
		} 
		dp[0][i] = initVal
	}
	
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
