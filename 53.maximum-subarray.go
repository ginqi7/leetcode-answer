package main
func maxSubArray(nums []int) int {
    n := len(nums)
	if n == 0 {
		return 0
	}
	pre := nums[0]
	ans := pre
	for i := 1; i < n; i++ {
		var current int
		if pre < 0 {
			current = nums[i]
		} else {
			current = pre+nums[i]
		}
		pre = current
		if ans < current {
			ans = current
		}
	}
	return ans
}
