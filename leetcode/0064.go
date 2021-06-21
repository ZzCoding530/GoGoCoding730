package leetcode

func minPathSum(grid [][]int) int {
	if grid == nil {
		return 0
	}

	// 初始化二维数组，记得每个维度都要手动make
	dp := make([][]int, len(grid))
	for k, _ := range grid {
		dp[k] = make([]int, len(grid[k]))
	}

	dp[0][0] = grid[0][0]

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i != 0 && j != 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i == 0 && j != 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i != 0 && j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
