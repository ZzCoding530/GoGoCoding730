package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var dummyHead *ListNode = &ListNode{}
	dummyHead.Next = head
	nowNode := dummyHead

	step := right - left

	for left > 1 {
		nowNode = nowNode.Next
		left--
	}

	//nowNode目前在left-1的位置上
	//现在 nowNode在第m-1位置上， curr在m位置上
	var prev *ListNode = nil
	var now *ListNode = nowNode.Next

	for step >= 0 {
		temp := now.Next
		now.Next = prev

		prev = now
		now = temp

		step--
	}

	nowNode.Next.Next = now
	nowNode.Next = prev

	return dummyHead.Next

}
