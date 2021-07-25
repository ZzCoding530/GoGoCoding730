package finalShoot

/*
	双指针，滑动窗口往后滑，留最小的窗口长
*/
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	ans := -1
	start, end := 0, 0
	sum := 0
	for end < n { //右边界条件
		sum += nums[end]    //更新sum值
		for sum >= target { //如果当前sum >= target
			ans = min(ans, end-start+1) //更新ans值
			sum -= nums[start]          //试着把左侧窗口往右缩，缩小窗口后记得把sum值也更新了
			start++
		}
		end++ //总体上右边界一直往后走
	}

	return ans
}
