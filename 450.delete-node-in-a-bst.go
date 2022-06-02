package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
		return root
	}
	if root.Val == key {
		left := root.Left
		right := root.Right
		if left == nil {
			return right
		}
		iterator := left
		for ; iterator.Right != nil; iterator = iterator.Right {}
		iterator.Right = right
		return left
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
