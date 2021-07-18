package finalShoot

/*
	dp 跟001题一样
*/
func countSquares(matrix [][]int) int {
	var ans = 0
	dp := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(dp[0]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			} else if matrix[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
			ans += dp[i][j]

		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
