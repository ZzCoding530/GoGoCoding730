package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	start := 0
	end := 0
	var appeared [256]int
	for i := 0; i < len(appeared); i++ {
		appeared[i] = -1
	}
	var result = 0
	for end < len(s) {

		//如何表示一个字符，用当前字符-'a'，得到一个数字，来表示
		//index 一开始都被设置成-1 如果当前字符上一次出现的位置不是-1，说明出现过了
		//如果上次出现的位置index 在start之后，那就说明重复了，
		//start就要被挪到index+1了，这样start和end之间暂时又没有重复元素了
		if idx := appeared[s[end]-'a']; idx != -1 && idx >= start {
			start = idx + 1
		}
		appeared[s[end]-'a'] = end
		result = max(result, end-start+1)
		end++
	}
	return result

}
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
