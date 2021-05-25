package leetcode

import "../structures"

type Node = structures.Node

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	/*
		原来的链表为A-B-C-D-E
		依次复制链表变为 A-a-B-b-C-c-D-d-E-e
	*/
	ptr := head
	for ptr != nil {

		var newNode *Node = &Node{Val: ptr.Val}
		newNode.Next = ptr.Next
		ptr.Next = newNode
		ptr = newNode.Next

	}

	/*
		依次给小写的Node的Random指针赋值
	*/
	ptr = head
	for ptr != nil {
		if ptr.Random == nil {
			ptr.Next.Random = nil
		} else {
			ptr.Next.Random = ptr.Random.Next
		}
		ptr = ptr.Next.Next

	}

	/*
		把A-a-B-b-C-c-D-d-E-e 拆成 a-b-c-d-e
	*/
	oldPtr := head
	newPtr := head.Next
	resultHead := head.Next

	for oldPtr != nil {
		oldPtr.Next = oldPtr.Next.Next
		if newPtr.Next != nil {
			newPtr.Next = newPtr.Next.Next
		} else {
			newPtr.Next = nil
		}

		oldPtr = oldPtr.Next
		newPtr = newPtr.Next
	}

	return resultHead
}
