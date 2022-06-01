package main
/**
 Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if (len(preorder) == 0) {
		return nil
	}
	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(preorder)-1)
}

func buildTreeHelper(preorder []int, preBegin, preEnd int, inorder []int, inBegin, inEnd int) *TreeNode {
	if (preBegin > preEnd) {
		return nil
	}
	root:= &TreeNode{preorder[preBegin], nil, nil}
	if preBegin == preEnd {
		return root
	}
	idx := findIdx(inorder, root.Val)
	leftSize := idx - inBegin
	root.Left = buildTreeHelper(preorder, preBegin+1, preBegin+leftSize, inorder, inBegin, idx-1)
	root.Right = buildTreeHelper(preorder, preBegin+leftSize + 1, preEnd, inorder, idx+1, inEnd)
	return root
}

func findIdx(arry[]int, a int) int {
	for i:=0; i <len(arry) ;i++ {
		if arry[i] == a {
			return i
		}
	}
	return 0
}
