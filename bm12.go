package main

func SortInList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left := head
	mid := head.Next
	right := head.Next.Next
	for right != nil && right.Next != nil {
		left = left.Next
		mid = mid.Next
		right = right.Next.Next
	}
	left.Next = nil
	return mergeTwo(SortInList(head), SortInList(mid))
}

func mergeTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
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
