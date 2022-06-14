package main
/**
 Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
 */
func removeElements(head *ListNode, val int) *ListNode {
    current := head
	root := &ListNode{0, head}
	pre := root 
	for current != nil {
		if current.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = current
		}
		current = current.Next
	}
	return root.Next
}
