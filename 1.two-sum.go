package main
func twoSum(nums []int, target int) []int {
	index_map := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := index_map[nums[i]]; ok {
			return []int{val, i}
		} else {
			index_map[target-nums[i]] = i;
		}
	}
	return []int{0, 0}
}
