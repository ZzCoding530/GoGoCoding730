package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummyHead *ListNode = &ListNode{Next: head}

	//这里，如果slow和fast都从head开始，那slow会停在倒数第K个节点
	//为了方便删除，我们需要slow停在倒数第K个节点的前一个节点
	//所以让slow的出发位置就往前一位，head的前驱节点就是dummyHead
	var slow *ListNode = dummyHead
	var fast *ListNode = head

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummyHead.Next

}
