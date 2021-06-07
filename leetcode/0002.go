package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	overTen := 0
	root := &ListNode{}
	ans := root

	for l1 != nil || l2 != nil {
		var v1, v2 int

		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}

		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}

		sum := (v1 + v2 + overTen) % 10
		overTen = (v1 + v2 + overTen) / 10

		root.Next = &ListNode{Val: sum}
		root = root.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	if overTen != 0 {
		root.Next = &ListNode{Val: overTen}
	}
	return ans.Next
}
