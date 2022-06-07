package main
func search(nums []int, target int) int {
    left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		val := nums[mid]
		if val == target {
			return mid
		} else if val > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
