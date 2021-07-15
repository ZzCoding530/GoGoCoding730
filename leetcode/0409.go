package leetcode

func longestPalindrome(s string) int {
	//用map记录每个字符出现的次数
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	var ans int
	for _, count := range m {
		//除以二是为了求这个字符有多少对，有几对，能够成的字符串长度就是 几对*2，这里面有个int向下取整所以这么写
		ans += count / 2 * 2
	}
	// 如果最终的长度小于原字符串的长度，说明里面某个字符出现了奇数次，那么那个字符可以放在回文串的中间，所以额外再加一。
	if ans < len(s) {
		return ans + 1
	}
	return ans
}
