package main

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	length := 0
	cur := pHead
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length < k {
		return nil
	}
	node := pHead
	for i := 0; i < k; i++ {
		node = node.Next
	}
	cur = pHead
	for node != nil {
		cur = cur.Next
		node = node.Next
	}
	return cur
}
