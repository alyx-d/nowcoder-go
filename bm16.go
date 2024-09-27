package main

func DeleteDuplicates2(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	cur := dummy
	for cur != nil && cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Val
			for cur.Next != nil && temp == cur.Next.Val {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
