package main
import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := [][]int{}
	for i:= 0 ; i < n; i++ {
		for j:=i+1; j < n; j++ {
			k := j + 1
			v := n - 1
			for k < v {
				sum := nums[i] + nums[j] + nums[k] + nums[v]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[v]})
					for ; k < n-1 && nums[k] == nums[k+1]; k++ {}
					for ; v > 0  && nums[v] == nums[v-1]; v-- {}
					k++
					v--
				} else if sum > target {
					for ; v > 0  && nums[v] == nums[v-1]; v-- {}
					v--
				} else {
					for ; k < n-1 && nums[k] == nums[k+1]; k++ {}
					k++
				}
			}
			for ; j < n-1 && nums[j] == nums[j+1]; j++ {}
		}
		for ; i < n-1 && nums[i] == nums[i+1]; i++ {}
	}
	return ans
    
}
