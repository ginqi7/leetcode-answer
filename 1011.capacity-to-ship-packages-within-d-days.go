package main
func shipWithinDays(weights []int, days int) int {
    left :=  Max(weights)
	right := Sum(weights)
	for left < right {
		mid := left + (right - left) / 2
		if f(weights, mid) > days {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left

}

//
func f(weights []int, capacity int) int {
	days := 1
	remaind := capacity
	for _, val := range weights {
		if remaind >= val {
			remaind -= val
		} else {
			days++
			remaind = capacity - val
		}
	}
	return days
}

func Max(weights []int) int{
	max := 0
	for _, val := range weights {
		if max < val {
			max = val
		}
	}
	return max
}


func Sum(weights []int) int{
	sum := 0
	for _, val := range weights {
		sum += val
	}
	return sum
}
