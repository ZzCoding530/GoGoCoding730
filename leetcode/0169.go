package leetcode

func majorityElement(nums []int) int {
	vote := 0
	x := nums[0]
	for _, num := range nums {
		// 票数抵消
		if vote == 0 {
			// 假设当前数为众数
			x = num
		}
		if num == x {
			// 投正票
			vote++
		} else {
			// 投反票
			vote--
		}
	}
	return x
}
