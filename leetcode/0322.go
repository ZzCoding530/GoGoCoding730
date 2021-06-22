package leetcode

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, amount+1)

	for i := 0; i <= amount; i++ {
		dp[i] = max
	}

	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] == max {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
