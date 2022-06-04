package main
func coinChange(coins []int, amount int) int {
    //  dp(n) :
	// 0 , n= 0
	// -1 , n < 0
	// min{dp(n-coid) + 1 | coin in coins}
	
	dp := []int{}
	for i:=0 ; i < amount+1; i++ {
		dp = append(dp, amount+1)
		if i == 0 {
			dp[0] = 0
		}
		for j := 0; j < len(coins); j++ {
			if i - coins[j] >= 0 {
				dp[i] = Min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	
	return dp[amount]
}

func Min(a, b int) int{
	if a > b {
		return b
	}
	return a
}

