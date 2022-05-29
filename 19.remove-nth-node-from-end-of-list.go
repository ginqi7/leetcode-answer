package main

// Definition for singly-linked list.
  // type ListNode struct {
  //     Val int
  //     Next *ListNode
  // }
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := ListNode{0, head} 
	pointA := &root
	pointB := &root 

	for ; n >= 0 ; n-- {
		if pointB == nil {
			return nil
		}
		pointB = pointB.Next
	}

	for ; pointB != nil; pointB = pointB.Next {
		pointA = pointA.Next
	}
	if pointA.Next != nil {
		pointA.Next = pointA.Next.Next
	}
	return root.Next 
}
