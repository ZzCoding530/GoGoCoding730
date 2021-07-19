package finalShoot

/*
	思路不难，但保留dummyHead很关键
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var dummyHead *ListNode = &ListNode{}
	dummyHead.Next = head
	nowNode := dummyHead

	step := right - left

	//假设第一个node的index为1，则nowNode现在停在 left-1 位置
	for left > 1 {
		nowNode = nowNode.Next
		left--
	}

	var temp *ListNode
	var prev *ListNode
	now := nowNode.Next

	for step >= 0 {
		temp = now.Next

		now.Next = prev
		prev = now
		now = temp

		step--
	}

	nowNode.Next.Next = now
	nowNode.Next = prev

	return dummyHead.Next // 一定要注意返回dummyHead.Next 不然链表从第一个开始反转（left=1）时候就错了
}
