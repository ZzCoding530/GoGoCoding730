package leetcode

/*
	有效括号 不难想出来是用单调栈
*/
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 { //个数不是偶数那就是有多余的 就不对呗
		return false
	}

	pairs := map[byte]byte{ //建一个map对应括号
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte //栈

	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { //这里用字符大于0判断map里有没有对应的字符
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { //false的情况：1.栈空了  2.栈顶元素和当前的对不上
				return false
			}

			stack = stack[:len(stack)-1] //如果能对上，那就把栈顶元素出栈，相当于抵消了
		} else { //没有对应字符那就把当前字符加进去
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
