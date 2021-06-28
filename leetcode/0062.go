package leetcode

func uniquePaths(m, n int) int {
	dp := make([][]int, m) //dp数组，机器人出发那点的横竖两条边都为1，都是直线一种方式到达

	//初始化二维数组记住了这么做
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1 //竖着的置为1
	}
	for j := 0; j < n; j++ { //横着的置为1
		dp[0][j] = 1
	}

	//dp
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1] //返回最后的角
}
