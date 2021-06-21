package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	queue := []int{} //双端队列，存的是index下标，不是数本身
	res := []int{}   //返回用的

	for i := 0; i < len(nums); i++ {

		//在放元素或者不放之前先判断，要是待放元素比目前队列尾的元素大，那就把队尾元素扔掉，注意是所有的比待放入元素小的，用for循环都扔掉
		for i > 0 && len(queue) > 0 && nums[queue[len(queue)-1]] < nums[i] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i) //把这个元素下标放进去

		//当滑动窗口形成即i>=k,并且队首值已经不在窗口中时候，要把队列最左边的丢掉
		//nums[i]是窗口最右侧那个元素，nums[i-k]是窗口最左侧的左边的那个元素，即左窗框外边的元素，是在窗口外面的
		//如果队首元素是这个，那就得扔出去了
		if queue[0] < i-k+1 {
			queue = queue[1:]
		}

		if i+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}

	return res

}
