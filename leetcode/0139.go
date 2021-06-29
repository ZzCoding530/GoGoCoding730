package leetcode

func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool) //新建一个map保存wordDict里的单词
	for _, w := range wordDict {         //单词都放进去
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1) //dp数组		dp[i]表示s中的前i位是否可以用wordDict中的单词表示
	dp[0] = true                 //空字符是可以被表示的
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] { //如果前j位都可以被表示	&&		字符串s的j到i位切出来的字符串正好是wordDict里的单词，
				dp[i] = true // 那就说明0-j前，j-i都是true，0-i位置就是true了
				break
			}
		}
	}
	return dp[len(s)]
}
