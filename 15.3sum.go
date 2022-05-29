package main
import (
	"sort"
)
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < n; i++ {
		j := i + 1
		k := n - 1
		for j < n && j < k {
			if nums[i] + nums[j] + nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for ; k != 0 && nums[k] == nums[k-1] ; k-- {}
				for ; j != n-1 && nums[j] == nums[j+1] ; j++ {}
				k--
				j++
			} else if nums[j] + nums[k] > -1 * nums[i] {
				for ; k != 0 && nums[k] == nums[k-1] ; k-- {}
				k--
			} else {
				for ; j != n-1 && nums[j] == nums[j+1] ; j++ {}
				j++
			}
		}
		for ; i != n-1 && nums[i] == nums[i+1] ; i++ {}
	}
	return ans
}


