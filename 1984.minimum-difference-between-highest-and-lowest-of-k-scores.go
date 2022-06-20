package main
func minimumDifference(nums []int, k int) int {
    sort.Ints(nums)
	n := len(nums)
	if n == 0 || n == 1 || k == 1{
		return 0
	}
	minDiff := math.MaxInt
	for i := 0; i < n - k + 1 ; i++ {
		j := i + k - 1
		if minDiff > nums[j] - nums[i] {
			minDiff = nums[j] - nums[i]
		}
	}
	return minDiff
}
