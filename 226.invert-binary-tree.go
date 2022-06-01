package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	right := root.Right
	root.Right = left
	root.Left = right
	invertTree(left)
	invertTree(right)
    return root
}
