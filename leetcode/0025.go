package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {

	//先找到第kth个node
	var kthNode *ListNode = head

	for i := 0; i < k; i++ {
		if kthNode == nil {
			return head
		}
		kthNode = kthNode.Next
	}

	newHead := reverse(head, kthNode)
	head.Next = reverseKGroup(kthNode, k)

	return newHead
}

func reverse(node1 *ListNode, node2 *ListNode) *ListNode {
	var prev *ListNode = nil
	now := node1
	for now != node2 {
		temp := now.Next
		now.Next = prev

		prev = now
		now = temp
	}

	return prev
}
