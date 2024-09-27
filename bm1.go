package main

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := ReverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return next
}

func ReverseList2(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}
