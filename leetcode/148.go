package leetcode

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := 0 //这段数一下链表的长度
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummyHead := &ListNode{Next: head}                        //哨兵节点
	for subLength := 1; subLength < length; subLength <<= 1 { // subLength是子串的长度 每一轮不断乘2
		tempDummyHead, curNodePtr := dummyHead, dummyHead.Next //tempDummyHead是在每个切分链表的循环里的小哨兵节点， curNodePtr是当前节点的指针
		for curNodePtr != nil {
			head1 := curNodePtr                                        //要按照subLength拆分的第一个链表头节点
			for i := 1; i < subLength && curNodePtr.Next != nil; i++ { //curNodePtr往后走，走subLength长度
				curNodePtr = curNodePtr.Next
			}

			head2 := curNodePtr.Next                                                        //curNodePtr的下一位就是 要按照subLength拆分的第二个链表的头节点
			curNodePtr.Next = nil                                                           //这句是让curNodePtr这个节点指向nil，  也就是	把链表从curNodePtr位置切断，
			curNodePtr = head2                                                              //curNodePtr指针挪到head2 即第二个链表的头节点
			for i := 1; i < subLength && curNodePtr != nil && curNodePtr.Next != nil; i++ { //这段找一下第二个链表的尾巴在哪
				curNodePtr = curNodePtr.Next
			}

			var next *ListNode //再次切断链表 这次是在第二条链表的尾部切断
			if curNodePtr != nil {
				next = curNodePtr.Next //这个next 记录了第二个链表切断后，剩下的链表头节点是哪
				curNodePtr.Next = nil  //这一句是切断操作
			}

			tempDummyHead.Next = merge(head1, head2) //合并两个链表，连在一开始设置的哨兵节点的后面

			for tempDummyHead.Next != nil { //把这轮拆分链表的小哨兵节点往后挪，一直到最尾部
				tempDummyHead = tempDummyHead.Next
			}
			curNodePtr = next //把指针指向刚才切完两个链表以后，剩下的链表的头节点
		}
	}
	return dummyHead.Next
}

/*
	这一段就是21题的合并两个有序链表
*/
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
