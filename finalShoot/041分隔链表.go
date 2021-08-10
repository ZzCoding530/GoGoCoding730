package finalShoot

/*
	两个链表，一个存小于x的，一个剩下的的
*/
func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}

	//两个链表构造好以后，连起来
	large.Next = nil
	small.Next = largeHead.Next
	return smallHead.Next
}
