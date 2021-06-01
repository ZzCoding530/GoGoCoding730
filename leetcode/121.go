package leetcode

/*
1、长度小于2， 直接返回0
2、小于最小值，设置为最小值
3、减去最小值大于最大收益，赋值最大收益
*/

func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	min, ans := prices[0], 0
	for i := 1; i < n; i++ {
		if tmp := prices[i] - min; tmp > ans {
			ans = tmp
		} else if prices[i] < min {
			min = prices[i]
		}
	}
	return ans
}
