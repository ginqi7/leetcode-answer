package main
import (
	"math"
)
func reverse(x int) int {
	if x < 0 {
		return -1 * reverseHelper(-1 * x, 0) 
	}
	return reverseHelper(x, 0)
	
}

func reverseHelper(x, last int) int {
	last = last * 10 + x % 10
	x = x / 10
	if last > math.MaxInt32 {
		return 0 
	}
	if x == 0 {
		return last
	}
	return reverseHelper(x, last)
	
}
