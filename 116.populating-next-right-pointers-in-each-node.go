package main 
// Definition for a Node.
  // type Node struct {
  //     Val int
  //     Left *Node
  //     Right *Node
  //     Next *Node
  // }
 

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode (root1, root2 *Node){
	if root1 == nil || root2 == nil {
		return 
	}
	root1.Next = root2
	connectTwoNode(root1.Left, root1.Right)
	connectTwoNode(root2.Left, root2.Right)
	connectTwoNode(root1.Right, root2.Left)
	
}

