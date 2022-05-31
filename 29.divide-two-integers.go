package main
import "math"
func divide(dividend int, divisor int) int {
    ans := dividend /divisor
	if (ans > math.MaxInt32) {
		return math.MaxInt32
	}
	return ans
}
