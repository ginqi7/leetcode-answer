package main
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n < 2 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-2] + cost[i-2], dp[i-1] + cost[i-1])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
