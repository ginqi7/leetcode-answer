package main
func calculateMinimumHP(dungeon [][]int) int {
    m := len(dungeon)
	if m == 0 {
		return 0
	}
	n := len(dungeon[0])
	
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = getLife(dungeon[i][j])
			} else if i == m-1 {
				dp[i][j] = getLifes(dp[i][j+1] , dungeon[i][j])
			} else if j == n-1 {
				dp[i][j] = getLifes(dp[i+1][j] , dungeon[i][j])
			} else {
				dp[i][j] = getLifes(Min(dp[i+1][j], dp[i][j+1]) , dungeon[i][j])
			}
		}
	}
	//fmt.Printf("%+v\n", dp) // output for debug

	return dp[0][0]

}

func getLifes(a, b int) int {
	if b - a < 0 {
		return -1 * (b-a)
	}
	return 1
}

//
func getLife(a int) int {
	if a > 0 {
		return 1
	} 
	return -1 * a + 1
}

//
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
