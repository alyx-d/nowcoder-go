package main

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	dummy := new(ListNode)
	l1 := pHead1
	l2 := pHead2
	cur := dummy
	for l1 != nil && l2 != nil {
		l := new(ListNode)
		if l1.Val < l2.Val {
			l.Val = l1.Val
			l1 = l1.Next
		} else {
			l.Val = l2.Val
			l2 = l2.Next
		}
		cur.Next = l
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next
}
