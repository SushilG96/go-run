/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

func reverseList(head *ListNode) *ListNode {
	cur := head
	tmp := []int{}
	for cur != nil {
		tmp = append(tmp, cur.Val)
		cur = cur.Next
	}
	cur1 := head
	for i := len(tmp) - 1; i >= 0; i-- {
		cur1.Val = tmp[i]
		cur1 = cur1.Next
	}
	return head
}
