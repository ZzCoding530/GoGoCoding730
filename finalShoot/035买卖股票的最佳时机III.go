package finalShoot

/*
	buy1,sell1,buy2,sell2 四个值分别记录第一次买卖和第二次买卖时候手里剩的钱
*/
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])       //如果不买，那就还是这些钱，如果买，就减去当前股票价格
		sell1 = max(sell1, buy1+prices[i]) //不卖就这点钱，卖了就挣到当天的股票钱
		buy2 = max(buy2, sell1-prices[i])  //要么不买，要么就第一次卖了以后的钱买今天的股票
		sell2 = max(sell2, buy2+prices[i]) //要么不卖，要么第二次买完赚到今天的股票价格
	}

	return sell2
}
