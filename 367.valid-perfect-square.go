package main
func isPerfectSquare(num int) bool {
    if num == 0 || num == 1 {
		return true
	}
	
	left := 0
	right := num
	
	for left < right {
		mid := left + (right - left) / 2
		val := mid * mid
		if val == num {
			return true
		} else if val > num {
			right = mid 
		} else {
			left = mid + 1
		}
	}
	
	return false
}
