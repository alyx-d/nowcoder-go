package main

func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	d := new(ListNode)
	d.Next = head
	pre := d
	cur := head
	for i := 1; i < m; i++ {
		pre = cur
		cur = cur.Next
	}
	for i := m; i < n; i++ {
		temp := cur.Next
		cur.Next = temp.Next
		temp.Next = pre.Next
		pre.Next = temp
	}
	return d.Next
}
