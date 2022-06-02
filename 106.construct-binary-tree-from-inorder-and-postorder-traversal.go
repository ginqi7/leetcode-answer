package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
		return nil
	}
	return buildTreeHelper(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}


func buildTreeHelper(inorder []int, inBegin, inEnd int, postorder []int, postBegin, postEnd int) *TreeNode {
	if (postEnd < postBegin) {
		return nil
	}
	val := postorder[postEnd]
	root := &TreeNode{val, nil, nil}
	if (postEnd == postBegin) {
		return root
	}
	idx := findIdx(inorder, val, inBegin, inEnd)
	leftSize := idx - inBegin
	root.Left = buildTreeHelper(inorder, inBegin, idx-1, postorder, postBegin, postBegin+leftSize-1)
	root.Right = buildTreeHelper(inorder, idx+1, inEnd, postorder, postBegin + leftSize, postEnd-1)
	return root
	
}

func findIdx(inorder []int, val int, inBegin, inEnd int) int {
	for i := inBegin; i <= inEnd; i++ {
		if inorder[i] == val {
			return i
		}
	}
	return inBegin
}
