package main
/**
 * Definition for singly-linked list.
 *  type ListNode struct {
 *      Val int
 *      Next *ListNode
 *  }
 **/ 

 
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
    root := &ListNode{0, head}
	pointA := root
	pointB := root
	pointA = head.Next
	for pointA != nil {
		pointBNext := pointB.Next
		pointANext := pointA.Next
		pointB.Next = pointA
		pointA.Next = pointBNext
		pointBNext.Next = pointANext
		pointB = pointBNext
		if pointANext != nil {
			pointA = pointANext.Next
		} else {
			pointA = nil
		}
	}
	return root.Next
}


