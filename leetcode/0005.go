package leetcode

/*
	动态规划dp
	dp[l][r]表示从l到r这段是不是回文
*/
func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	dp := make([][]bool, l)

	ans := 1
	resultStart := 0

	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
	}

	for j := 1; j < l; j++ {

		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			}

			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}

			}

			if dp[i][j] && j-i+1 > ans {
				ans = j - i + 1
				resultStart = i
			}

		}
	}

	return s[resultStart : resultStart+ans]
}
