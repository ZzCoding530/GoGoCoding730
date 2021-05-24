package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	headNext := head.Next
	head.Next = swapPairs(headNext.Next)
	headNext.Next = head
	return headNext
}
