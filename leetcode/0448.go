package leetcode

/*
	跟442思路一模一样
*/
func findDisappearedNumbers(nums []int) []int {
	res := []int{}

	for i := 0; i < len(nums); i++ {
		nums[abs(nums[i])-1] = -abs(nums[abs(nums[i])-1])
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
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
