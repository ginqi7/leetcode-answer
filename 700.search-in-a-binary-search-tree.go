package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
		return nil
	}
	cVal := root.Val
	if  cVal == val {
		return root
	} else if cVal > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
