package main
// Definition for singly-linked list.
//  type ListNode struct {
//      Val int
//      Next *ListNode
//  }
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c, val := sum(l1.Val, l2.Val, 0)
	root := ListNode{Val: val}
	pre := &root
	for getNext(l1) != nil || getNext(l2) != nil {
		l1 = getNext(l1)
		l2 = getNext(l2)
		c, val = sum(getValue(l1), getValue(l2), c)
		node := ListNode{Val: val}
		pre.Next = &node
		pre = pre.Next
	}
	if c != 0 {
		node := ListNode{Val: c}
		pre.Next = &node
	}
	return &root
}

func sum(n1 int, n2 int, n3 int) (int, int) {
	sum := n1 + n2 + n3
	return sum / 10, sum % 10
}

func getNext(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}
	return nil
}

func getValue(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
}
