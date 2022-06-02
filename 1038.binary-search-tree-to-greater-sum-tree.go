package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
var sum int
func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
    bstToGstHelper(root)
	return root
}

func bstToGstHelper(root *TreeNode) {
	if root == nil {
		return
	}
	bstToGstHelper(root.Right)
	sum += root.Val
	root.Val = sum
	bstToGstHelper(root.Left)
}
