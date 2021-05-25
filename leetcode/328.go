package leetcode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	evenHead := head.Next
	oddPtr := head
	evenPtr := evenHead

	for evenPtr != nil && evenPtr.Next != nil {
		oddPtr.Next = evenPtr.Next
		oddPtr = oddPtr.Next

		evenPtr.Next = oddPtr.Next
		evenPtr = evenPtr.Next
	}

	oddPtr.Next = evenHead

	return head
}
