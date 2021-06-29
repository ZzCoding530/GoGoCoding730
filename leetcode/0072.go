package leetcode

func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	//初始化了一个二维数组  dp[i][j]表示从 word1[0:i] ➡️ word2[0:j] 需要多少步
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	//初始条件：
	for i := 0; i < len1+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < len2+1; j++ {
		dp[0][j] = j
	}

	//进行运算
	for a := 1; a <= len1; a++ {
		for b := 1; b <= len2; b++ {

			if word1[a-1] == word2[b-1] { //如果相等，那就不用操作这一位
				dp[a][b] = dp[a-1][b-1]
			} else { //不相等的情况
				dp[a][b] = min(min(dp[a][b-1]+1, dp[a-1][b]+1), dp[a-1][b-1]+1)
				//dp[a][b-1]+1： 	word2比1短一位那就增加一位
				//dp[a-1][b]+1： 	word2比1长一位那就减少一位
				//dp[a-1][b-1]+1：	word2跟1一样长那就替换一位
				//取操作数最小的存进去
			}

		}

	}

	return dp[len1][len2]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
