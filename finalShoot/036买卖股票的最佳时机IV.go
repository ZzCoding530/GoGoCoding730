package finalShoot

import "math"

func maxProfit(k int, prices []int) int {
	if len(prices) < 2 || k == 0 {
		return 0
	}
	var res int
	k = min(k, len(prices)/2)
	// buy[i][j],第i天，进行j笔交易，
	// buy[i][j] = max(buy[i-1][j],sell[i-1][j-1]-prices[i-1])
	// sell[i][j],第i天，进行j笔交易，不持有股票的利润
	// sell[i][j] = max(sell[i-1][j],buy[i-1][j]+prices[i-1])
	buy := make([][]int, len(prices)+1)
	sell := make([][]int, len(prices)+1)
	for i, _ := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	//上面是初始化数组

	buy[1][1] = -prices[0]
	sell[1][1] = 0
	for i := 2; i <= k; i++ {
		// 在第一天不可能完成两笔以上交易
		buy[1][i], sell[1][i] = math.MinInt32, math.MinInt32
	}
	for i := 2; i <= len(prices); i++ {
		buy[i][1] = max(buy[i-1][1], sell[i-1][0]-prices[i-1])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j-1]-prices[i-1])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j]+prices[i-1])
		}
	}
	for i := 1; i <= len(prices); i++ {
		for j := 1; j <= k; j++ {
			res = max(res, sell[i][j])
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
