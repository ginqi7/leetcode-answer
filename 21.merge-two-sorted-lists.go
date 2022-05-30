/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	ans := &ListNode{0, nil}
	iterator := ans
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			iterator.Next = list1
			list1 = list1.Next
		} else {
			iterator.Next = list2
			list2 = list2.Next
		}
		iterator = iterator.Next
	}
	if list1 != nil {
		iterator.Next = list1
	} else if list2 != nil {
		iterator.Next = list2
	}
	 return ans.Next
	
}
