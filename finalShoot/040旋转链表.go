package finalShoot

/*
	把链表构造成环形链表再拆开
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	size := 0
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		size++
	}

	size++          //目前cur停在最后一个节点，所以还得再加1
	cur.Next = head // 将尾巴指向头部
	cur = head      // 重新指向头部
	k = k % size    // 若k比长度大，那就套圈了
	for i := 1; i < size-k; i++ {
		cur = cur.Next
	}
	result := cur.Next
	cur.Next = nil
	return result
}
