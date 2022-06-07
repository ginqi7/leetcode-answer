package main
func minEatingSpeed(piles []int, h int) int {
    left := 1;
	right := maxK(piles)
	for left < right {
		mid := left + (right-left) / 2
		if f(piles, mid) > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func maxK(piles []int) int {
	k := 0
	for _, val := range piles {
		if val > k {
			k = val
		}
	}
	return k
}

//
func f(piles []int, k int) int {
	// Monotonically decreasing
	time := 0
	for _, val := range piles {
		time += val / k
		if val % k != 0 {
			time += 1
		}
	}
	return time
}
