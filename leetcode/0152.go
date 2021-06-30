package leetcode

/**
因为可能存在负数，最大值有可能是 负数*负数 得到的，所以要维护两个变量：最大值和最小值
*/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curMin, curMax := 1, 1
	res := nums[0]
	for _, num := range nums {
		preMax := curMax
		curMax = max(max(preMax*num, curMin*num), num)
		curMin = min(min(preMax*num, curMin*num), num)
		res = max(res, curMax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
