package finalShoot

//给你一个 m * n 的矩阵，矩阵中的元素不是 0 就是 1，请你统计并返回其中完全由 1 组成的 正方形 子矩阵的个数。

/*
	记住这种，求个数、求长度、求有几种方法几种路径、判断能不能到达、存在不存在，不需要你写出具体内容的，都往DP上想
*/
func maximalSquare(matrix [][]byte) int {
	var maxRes = 0
	dp := make([][]int, len(matrix)) //首先初始化二维数组，记得二维数组初始化要一层一层make
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ { //这里为了一会的转移方程好写，要把'1'位置的对应dp位置也置成1
		for j := 0; j < len(matrix[0]); j++ { //不然一会边缘不太好弄
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxRes = 1
			}
		}
	}

	for i := 1; i < len(dp); i++ { //这里从1开始，因为转移方程里有i-1，从0开始越界了
		for j := 1; j < len(dp[0]); j++ { //所以也要提前把边缘的数字填进去，就是刚才的两层循环做的事
			if matrix[i][j] == '1' { //这里用matrix和dp两个数组判断都一样
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1 //转移方程
				if dp[i][j] > maxRes {                                        //更新最大值
					maxRes = dp[i][j]
				}
			}
		}
	}

	return maxRes * maxRes //记得返回的是面积，dp存的是边长

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
