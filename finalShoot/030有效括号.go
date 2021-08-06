package finalShoot

/*
	使用map和栈
*/
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 { //奇数就显然是false
		return false
	}

	pairs := map[byte]byte{ //对应的map
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { //没匹配上就false
				return false
			}
			stack = stack[:len(stack)-1] //匹配上了就取出去栈顶的
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
