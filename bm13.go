package main

func IsPail(head *ListNode) bool {
	count := 0
	cur := head
	st := make([]*ListNode, 0)
	for cur != nil {
		count++
		st = append(st, cur)
		cur = cur.Next
	}
	cur = head
	for i := count - 1; i > count/2; i-- {
		if st[i : i+1][0].Val != cur.Val {
			return false
		}
		cur = cur.Next
	}
	return true
}
