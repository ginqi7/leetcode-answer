package main
func subarraySum(nums []int, k int) int {
	preSum := make(map[int]int)
	preSum[0] = 1
	
	ans := 0
	sum_0_i := 0
	for _, val := range nums {
		sum_0_i += val
		sum_0_j := sum_0_i - k
		val, ok := preSum[sum_0_j]
		if ok {
			ans += val
		}
		
		preSum[sum_0_i]++
	}
	return ans
}
