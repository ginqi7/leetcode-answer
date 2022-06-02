package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
func isValidBST(root *TreeNode) bool {
    return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(root, min, max *TreeNode) bool {
    if root == nil {
		return true
	}
	if (min != nil && min.Val >= root.Val) {
		return false
	}
	if (max != nil && max.Val <= root.Val) {
		return false 
	}
	return isValidBSTHelper(root.Left, min, root) &&
	isValidBSTHelper(root.Right, root, max)
}

