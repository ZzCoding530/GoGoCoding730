package leetcode

/*
	双指针，设置左边界和右边界，题目要求的是将奇数放左边，偶数放右边
	左指针遇到偶数就停止
	右指针遇到奇数停止
	在他们都停止的时候将左右指针交换一下子
*/
func exchange(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]%2 == 0 && nums[r]%2 != 0 {
			nums[l], nums[r] = nums[r], nums[l]
		}
		for l < len(nums) && nums[l]%2 != 0 {
			l++
		}
		for r >= 0 && nums[r]%2 == 0 {
			r--
		}

	}
	return nums
}
