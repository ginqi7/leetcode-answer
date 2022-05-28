package main
import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	/**
	  * using binary search to sovle this problem
	  * the time complexity is O(log(min(n ,m)))
	  * the search target is : 
	  * find four element, leftA, rightA, leftB and rightB, they satisfy:
	  * leftA <= rightB && leftB <= rightA  
	 **/
	n := len(nums1)
	m := len(nums2)
	if n > m {
		return findMedianSortedArrays(nums2, nums1)
	}

	theMid := (n + m + 1) / 2
	start := 0
	end := n

	for start <= end {
		mid := (start + end) / 2
		leftAsize := mid
		leftBsize := theMid - mid

		leftA := GetLeft(nums1, leftAsize) 
		rightA := GetRight(nums1, leftAsize, n)	
		leftB := GetLeft(nums2, leftBsize) 
		rightB := GetRight(nums2, leftBsize, m)	

		if leftA <= rightB && leftB <= rightA {
			if (m + n) % 2 == 0 {
				return (float64(Max(leftA, leftB)) + float64(Min(rightA, rightB))) / 2
			} else {
				return float64(Max(leftA, leftB))
			}
		} else if leftA > rightB {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0.0
}

func Min(a, b int) int {
	if a < b {
		return a;
	}
	return b;
}

func Max(a, b int) int {
	if a < b {
		return b;
	}
	return a;
}

func GetLeft(nums []int, size int) int {
	if (size > 0) {
		return nums[size-1]
	}
	return math.MinInt
}

func GetRight(nums []int, size int, len int) int {
	if size < len {
		return nums[size]
	}
	return math.MaxInt
}

