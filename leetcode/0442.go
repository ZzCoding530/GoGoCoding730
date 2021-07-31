package leetcode

func findDuplicates(nums []int) []int {
	res := []int{}
	n := len(nums)

	//有一个前提：数组里的数都是 1 - n 的
	for i := 0; i < n; i++ { // 比如现在index = i ，nums[i] = x , 那么让nums[x-1]置为负值； 如果已经是负数了，那说明 x 已经出现过一次了
		if nums[abs(nums[i])-1] <= 0 {
			res = append(res, abs(nums[i]))
		} else {
			nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
		}
	}

	return res

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
