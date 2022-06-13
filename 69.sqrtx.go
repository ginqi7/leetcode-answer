package main
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
    left := 1
	right := x
	for left < right {
		mid := left + (right - left) / 2
		val := mid * mid
		if val == x {
			return mid
		} else if val > x {
			right = mid 
		} else {
			left = mid + 1
		}
	}
	return right-1
}
