package main
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
    minLen := 0
	left := 0
	right := 0
	need := 0

	for right < n {
		need += nums[right]
		right++
		for need >= target {
			currentLen := right - left 
			if minLen == 0 || currentLen < minLen {
				minLen = currentLen
			}
			need -= nums[left]
			left ++
		}
	}
	return minLen
}
