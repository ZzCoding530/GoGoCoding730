package leetcode

func change(amount int, coins []int) int {
	dp := make([]int, amount+1) //dp数组
	dp[0] = 1                   //和为0只有一种

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = dp[i] + dp[i-coin]
		}
	}
	return dp[amount]
}
