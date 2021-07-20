package finalShoot

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = &ListNode{} //这里一定要new出来一个ListNode，不然会显示可能指向空指针
	var dummyHead *ListNode = head
	carry := 0

	for l1 != nil || l2 != nil { //这里最好写 或者 不要写 "l1!=nil&&l2!=nil"	不然碰到[9,9,9,9,9,9,9]，[9,9,9,9]，麻烦
		a := 0
		b := 0

		if l1 != nil {
			a = l1.Val
		}

		if l2 != nil {
			b = l2.Val
		}

		sum := (carry + a + b) % 10
		carry = (carry + a + b) / 10

		head.Next = &ListNode{Val: sum}
		head = head.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}
