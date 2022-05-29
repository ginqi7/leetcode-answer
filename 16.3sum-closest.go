package main
import (
	"sort"
	"math"
)
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closest := math.MaxInt
	closestSum := 0
	for i:=0; i < n-2; i++ {
		j := i+1
		k := n-1
		for j < k {
			sum := nums[i]+nums[j]+nums[k]
			if sum == target {
				return target
			}
			c := abs(sum-target)
			if (closest > c) {
				closest = c
				closestSum = sum
			}
			if sum > target {
				k--
			} else {
				j++
			}	
		}
	}
	return closestSum;
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}
