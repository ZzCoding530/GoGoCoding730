package finalShoot

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	自建堆实现吧
*/
func mergeKLists(lists []*ListNode) *ListNode {

}

type minHeap []*ListNode //别名

func (h minHeap) len() int           { return len(h) }           //长度就是数组长度
func (h minHeap) less(a, b int) bool { return a < b }            //比大小
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] } //交换
