package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	a := &ListNode{}
	b := false
	c := a
	for l1 != nil && l2 != nil {
		v := l1.Val + l2.Val
		if b {
			v = v + 1
		}
		if v >= 10 {
			v = v - 10
			b = true
		} else {
			b = false
		}

		a.Next = &ListNode{Val: v}
		a = a.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {

		if b {
			a.Next = &ListNode{Val: l1.Val+1}
			a = a.Next
			b = false
		}
		l1 = l1.Next
	}

	for l2 != nil {
		if b {
			a.Next = &ListNode{Val: l2.Val+1}
			a = a.Next
			b = false
		}
		l2 = l2.Next
	}
	if b {
		node := &ListNode{
			Val: 1,
		}
		a.Next = node
		a=a.Next
	}
	return c.Next
}
