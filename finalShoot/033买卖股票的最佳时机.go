package finalShoot

/*
	思路就是遍历，维护一个最小值min，不断用当前的计算ans并更新ans和min
*/
func maxProfit(prices []int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}

	min := prices[0]
	ans := 0

	for i := 0; i < n; i++ {
		if temp := prices[i]; temp-min > ans {
			ans = temp - min
		} else if prices[i] < min {
			min = prices[i]
		}
	}

	return ans
}
