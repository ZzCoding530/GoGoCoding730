package leetcode

func reorderList(head *ListNode) {

	if head == nil {
		return
	}

	//总体思路
	//1.先把链表分两半
	//2.后半段反转
	//3.把前半段和反转后的后半段，交替链接成新的

	midNode := findMiddleNode(head)
	var l1 *ListNode
	var l2 *ListNode
	l1 = head
	l2 = midNode.Next

	midNode.Next = nil
	l2 = reverseList(l2)

	mergeTwoList(l1, l2)

}

/*
 * 找中点 快慢指针
 */
func findMiddleNode(head *ListNode) *ListNode {
	var slow *ListNode = head
	var fast *ListNode = head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

/*
 * 反转链表
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

/*
 * 合并两个链表
 */
func mergeTwoList(l1 *ListNode, l2 *ListNode) {
	var temp1 *ListNode
	var temp2 *ListNode

	for l1 != nil && l2 != nil {
		temp1 = l1.Next
		temp2 = l2.Next

		l1.Next = l2
		l1 = temp1

		l2.Next = l1
		l2 = temp2

	}
}
