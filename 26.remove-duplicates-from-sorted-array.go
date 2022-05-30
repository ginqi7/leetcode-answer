package main
func removeDuplicates(nums []int) int {
    count := 1
	if len(nums) == 0 {
		return 0
	}
	current := nums[0]
	for i := 1; i < len(nums) ; i++ {
		if current != nums[i] {
			current = nums[i]
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
