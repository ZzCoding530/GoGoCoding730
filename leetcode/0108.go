package leetcode

func rob(nums []int) int {
	//注意边界条件的处理
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	res := 0 //输出的结果
	l := len(nums)
	dp := [100]int{} //初始化一个数组

	//dp[i]表示走到第i家最多的钱是多少可能
	dp[0] = nums[0]               //到第一家时候 最多就是 偷了这家呗
	dp[1] = max(nums[1], nums[0]) //到第二家时候  比一下，要是前面那家多，那就不偷挨着这家了，不然偷这家

	for i := 2; i < l; i++ {
		//dp[i]就两种可能
		//1	不偷这家，那到dp[i]为止最多拿到的跟前一站一样
		//2 偷这家，那dp[i]就最多跟隔一家那里一样
		//取max
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		res = max(res, dp[i])
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
