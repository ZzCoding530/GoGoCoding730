package leetcode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var A *ListNode = headA
	var B *ListNode = headB

	for A != B {
		if A != nil {
			A = A.Next
		} else {
			A = headB
		}

		if B != nil {
			B = B.Next
		} else {
			B = headA
		}
	}

	return A

}
