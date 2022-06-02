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
func convertBST(root *TreeNode) *TreeNode {
	sum = 0
    sumBST(root)
	return root
}

func sumBST(root *TreeNode){
	if root == nil {
		return 
	}
	sumBST(root.Right)
	sum = sum + root.Val
	root.Val = sum
	sumBST(root.Left)
}
