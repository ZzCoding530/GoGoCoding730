package finalShoot

func reorderList(head *ListNode) {

	if head == nil {
		return
	}

	//总体思路
	//1.先把链表分两半
	//2.后半段反转
	//3.把前半段和反转后的后半段，交替链接成新的

	midNode := findMiddleNode(head) //返回的前半段链表尾部
	var l1 *ListNode
	var l2 *ListNode
	l1 = head         //前半段链表
	l2 = midNode.Next //后半段链表

	midNode.Next = nil   //截断连接
	l2 = reverseList(l2) //反转第二段链表

	mergeTwoList(l1, l2)

}

/*
 * 找中点 快慢指针 , for循环判断条件是 next 和 next.next 都不为nil	返回的node是前半段链表的尾部
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
 * 反转链表	记得返回prev 因为now已经在尾巴.next了，是nil
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
