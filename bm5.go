package main

func MergeKLists(lists []*ListNode) *ListNode {
	return partition(lists, 0, len(lists)-1)
}

func partition(lists []*ListNode, left int, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := (left + right) >> 1
	return merge(partition(lists, left, mid), partition(lists, mid+1, right))
}

func merge(l1 *ListNode, l2 *ListNode) *ListNode {
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
