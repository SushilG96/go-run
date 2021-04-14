package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	a1 := list1
	b1 := list2

	t := 0
	for cur1 != nil {
		if t == a-1 {
			a1 = cur1
		} else if t == b {
			b1 = cur1
		}
		t++
		cur1 = cur1.Next
	}
	a1.Next = cur2
	for cur2 != nil {
		if cur2.Next == nil {
			cur2.Next = b1.Next
			break
		}
		cur2 = cur2.Next
	}
	return list1
}
