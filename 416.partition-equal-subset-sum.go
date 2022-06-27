package main

func canPartition(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	sum := Sum(nums)
	if sum % 2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target + 1)
	
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
		}
	}
	return dp[target] == target
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func Sum(nums []int) int{
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += nums[i]
	}
	return ans
}
