package main
func maxArea(height []int) int {
	n := len(height)
	left := 0
	right := n - 1
	maxArea := 0
	for left < right {
		area := (right - left)  * Min(height[right], height[left])
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
