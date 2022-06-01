package main
import "math"
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	} 
    return constructMaximumBinaryTreeHelper(nums, 0, n-1)
}

func constructMaximumBinaryTreeHelper(nums []int, begin, end int) *TreeNode {
	if begin > end {
		return nil
	}
	if begin == end {
		return &TreeNode{nums[begin], nil, nil}
	}
	idx := findMaxIndex(nums, begin, end)
	left := constructMaximumBinaryTreeHelper(nums, begin, idx - 1)
	right := constructMaximumBinaryTreeHelper(nums, idx + 1, end)
	return &TreeNode{nums[idx], left, right}
}

func findMaxIndex(nums []int, begin, end int) int {
	maxVal:= math.MinInt
	idx := 0
	for i := begin; i <= end; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			idx = i
		}
	}
	return idx
}
