package leetcode

/*
	dp[i][j]表示 A[:i+1] 和 B[:j+1] 两个数组的公共子串
*/
func findLength(A []int, B []int) int {
	m, n := len(A), len(B)
	res := 0 //存结果的

	//这两步是初始化切片
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	//如果 A[i-1] != B[j-1]， 有 dp[i][j] = 0
	//如果 A[i-1] == B[j-1] ， 有 dp[i][j] = dp[i-1][j-1] + 1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > res { //更新长度
				res = dp[i][j]
			}
		}
	}
	return res
}
