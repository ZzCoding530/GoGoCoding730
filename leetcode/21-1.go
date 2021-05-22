package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	//！！！！下面这句非常非常重要！！！！一定要给哨兵节点做一个内存的初始化，不然没法往后指
	var dummyHead *ListNode = &ListNode{Val: 0, Next: nil} //假的头节点 真正的头节点是这个节点的下一位
	var nowNode = dummyHead                                // 当作指针往后走的节点

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			nowNode.Next = l1
			nowNode = nowNode.Next
			l1 = l1.Next
		} else {
			nowNode.Next = l2
			nowNode = nowNode.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		nowNode.Next = l2
	}

	if l2 == nil {
		nowNode.Next = l1
	}

	return dummyHead.Next
}
