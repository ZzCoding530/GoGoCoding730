package leetcode

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var fast *ListNode = head
	var slow *ListNode = head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
