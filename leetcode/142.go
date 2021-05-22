package leetcode

func detectCycle(head *ListNode) *ListNode {
	var slow *ListNode = head
	var fast *ListNode = head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			break
		}
	}

	//上面的步骤类似于141题
	// 如果fast真的走到了最后已经，那就是链表没环，返回nil
	if fast == nil || fast.Next == nil {
		return nil
	}

	//要搞清第二次相遇时候 两个指针都在环的入口处这一情况
	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
