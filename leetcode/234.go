package leetcode

func isPalindrome(head *ListNode) bool {
	middleNode := findMiddleNode(head)
	var newList *ListNode = reverse(middleNode)

	for head != nil && newList != nil {
		if newList.Val != head.Val {
			return false
		}

		newList = newList.Next
		head = head.Next
	}

	return true
}

func findMiddleNode(head *ListNode) *ListNode {
	var fast *ListNode = head
	var slow *ListNode = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	var now *ListNode = head
	for now != nil {
		temp := now.Next

		now.Next = prev

		prev = now
		now = temp
	}

	return prev
}
