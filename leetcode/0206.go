package leetcode

import "../structures"

type ListNode = structures.ListNode

/*
	206.反转链表
*/
func reverseList(head *ListNode) *ListNode {

	var prev *ListNode = nil
	var now *ListNode = head

	for now != nil {
		temp := now.Next

		now.Next = prev
		prev = now
		now = temp

	}

	return prev
}
