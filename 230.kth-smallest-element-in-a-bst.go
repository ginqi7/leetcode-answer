package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
var rank, ans int
func kthSmallest(root *TreeNode, k int) int {
	rank = 0
    kthSmallestHelper(root, k)
	return ans
}


func kthSmallestHelper(root *TreeNode, k int) {
	if root == nil {
		return 
	}
	kthSmallestHelper(root.Left, k);
	rank++
	if k == rank {
		ans = root.Val
	}
	kthSmallestHelper(root.Right, k)
}
