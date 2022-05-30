
package main

// Definition for singly-linked list.
  // type ListNode struct {
  //     Val int
  //     Next *ListNode
  // }
 
func mergeKLists(lists []*ListNode) *ListNode {
	for len(lists) > 1 {
		n := len(lists)
		newLists := []*ListNode{}
		for i := 0 ; i < n; i = i + 2 {
			mergeNode := sortMerge(getElement(lists, i), getElement(lists, i+1))
			if mergeNode != nil {
				newLists = append(newLists, mergeNode)
			}
		}
		lists = newLists
	}
	if (len(lists) == 0) {
		return nil
	}
	return lists[0]
}

func sortMerge(list1, list2 *ListNode) *ListNode {
	root := &ListNode{0, nil}
	iterator := root
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
	}
	if list2 != nil {
		iterator.Next = list2
	}
	return root.Next
}

func getElement(lists []*ListNode, idx int) *ListNode {
	if idx < len(lists) {
		return lists[idx]
	}
	return nil
}
