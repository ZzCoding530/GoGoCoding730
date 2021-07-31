package finalShoot

func heapSortArray(nums []int) []int {
	n := len(nums)
	buildHeap(nums, n) //建完堆以后

	for i := n - 1; i >= 0; i-- { // nums[0]是最大值，直接交换到末尾，对交换完的树再进行heapify调整，从尾巴到头部
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, i, 0)
	}

	return nums

}

/*
	也可以叫下沉sink过程，让小的往下沉，比如现在要对 index = i 的数据进行sink
*/
func heapify(nums []int, n, i int) {
	if i >= n {
		return
	}

	// index = i 的节点左右子节点分别为 2*i+1, 2*i+2
	child1 := 2*i + 1
	child2 := 2*i + 2
	max := i

	if child1 < n && nums[child1] > nums[max] {
		max = child1
	}
	if child2 < n && nums[child2] > nums[max] {
		max = child2
	} //到这步，max、child1、child2中的最大值存在了max里
	if max != i { //如果index = i的这个点不是这三个点中最大的，那这个点就得往下沉了，往下沉完势必破坏了下一层的结构，对下一层进行sink
		nums[max], nums[i] = nums[i], nums[max]
		heapify(nums, n, max)
	}
	//一直heapify到max 就是index为止，整个树就调整好了
}

/*
	建树的过程
*/
func buildHeap(nums []int, n int) {
	//这里 n/2 - 1 其实是 len(nums)/2-1 其实是最后一个元素的父节点的index，从最下层父节点开始，依次往上进行heapify
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
}
