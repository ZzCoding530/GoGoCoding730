package finalShoot

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	sum := []int{} //前缀和数组
	sum = append(sum, nums[0])
	for i := 1; i <= n; i++ { //填充前缀和数组
		sum = append(sum, sum[i-1]+nums[i-1])
	}
	queue := []int{}
	min := n + 1

	for i := 0; i < n; i++ {
		//当单调栈不为空    且  当前元素的前缀和 <= 栈顶元素的前缀和
		for len(queue) != 0 && sum[queue[len(queue)-1]] >= sum[i] {
			queue = queue[:len(queue)-1] //弹出栈顶元素
		}

		//若经过刚才的弹栈，还不为空 且 当前index 到 栈顶index的和>=k
		//其实意思就是。    试着把左侧窗口往右收， i 相当于是右侧窗口
		for len(queue) != 0 && sum[i]-sum[queue[len(queue)-1]] >= k {
			min = minInt(min, i-queue[len(queue)-1])
			queue = queue[1:]
		}
		queue = append(queue, i)
	}
	if min == n+1 {
		return -1
	}
	return min
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}
