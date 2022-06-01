package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
func flatten(root *TreeNode)  {
	if root == nil {
		return 
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	iterator := left
	if iterator == nil {
		return
	}
	for iterator.Right != nil {
		iterator = iterator.Right
	}
	iterator.Right = root.Right
	root.Left = nil
	root.Right = left


}


