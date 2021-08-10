package finalShoot

/*
	假设两个链表	长度分别为m,n
	假设两个链表	从头到相交处的距离为a,b
	假设两个链表	相交的长度为

	那么第一次相遇后走到尾部	链表A走了 a + c = m
							链表B走了 b + c = n

	如果再把两个指针放到对方的头部，那再相遇时候 走的距离分别为 a + c + b , b + c + a
	所以第二次相遇是停留在相交的节点
*/

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A := headA
	B := headB

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
