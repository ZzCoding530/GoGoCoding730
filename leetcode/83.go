package leetcode

func deleteDuplicates(head *ListNode) *ListNode {
	var now *ListNode = head

	for now != nil && now.Next != nil {
		if now.Val == now.Next.Val {
			now.Next = now.Next.Next
		} else {
			now = now.Next
		}
	}

	return head
}
