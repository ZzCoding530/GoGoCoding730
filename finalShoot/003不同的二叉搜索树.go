package finalShoot

/*
	跟我说的一样，求个数长度能否总数，不用列举的，都是DP
	但这题的转移方程需要看一下题解
	https://leetcode-cn.com/problems/unique-binary-search-trees/solution/hua-jie-suan-fa-96-bu-tong-de-er-cha-sou-suo-shu-b/
*/
func numTrees(n int) int {
	dp := make([]int, n+1)

	dp[1] = 1
	dp[0] = 1

	for i := 2; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]

}
