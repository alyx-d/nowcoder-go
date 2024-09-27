package main

func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	l1 := pHead1
	l2 := pHead2
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = pHead2
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = pHead1
		}
	}
	return l1
}
