package main
func searchRange(nums []int, target int) []int {
    return []int {searchLeft(nums, target), searchRight(nums, target)}
}

func searchLeft(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	
	if left == len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1

}

func searchRight(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	
	if left == 0 {
		return -1
	}
	if nums[left-1] == target {
		return left - 1
	}
	return -1
}
