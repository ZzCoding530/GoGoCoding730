package finalShoot

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	分治
*/
func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}

	if n == 1 {
		return lists[0]
	}

	return merge2list(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func merge2list(list1, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	head := dummyHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}

		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	}

	if list2 != nil {
		head.Next = list2
	}

	return dummyHead.Next
}
