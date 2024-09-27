package main

func AddInList(head1 *ListNode, head2 *ListNode) *ListNode {
	l1 := reverse(head1)
	l2 := reverse(head2)
	dummy := new(ListNode)
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		l := new(ListNode)
		l.Val = sum % 10
		carry = sum / 10
		cur.Next = l
		cur = cur.Next
	}
	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var newHead *ListNode = nil
	for head != nil {
		temp := head.Next
		head.Next = newHead
		newHead = head
		head = temp
	}
	return newHead
}
