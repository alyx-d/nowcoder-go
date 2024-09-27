package main

func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow := entry(pHead)
	if slow == nil {
		return nil
	}
	fast := pHead
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func entry(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}
	slow := pHead
	fast := pHead
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
