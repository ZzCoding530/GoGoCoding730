package leetcode

func longestValidParentheses(s string) int {
	maxAns := 0
	var stack []int
	stack = append(stack, -1) //先放一个-1，来处理第一个字符就是右括号的场合

	//遍历每个字符
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { //如果是左括号直接入栈（下标，index）
			stack = append(stack, i)
		} else { //如果是右括号
			stack = stack[:len(stack)-1] //先弹出栈顶元素
			if len(stack) == 0 {         //如果弹出后，栈空了，表示这个右括号没有匹配到左括号，要把这个右括号的index放进去
				stack = append(stack, i)
			} else { //如果弹出后还有，那长度就等于当前index - 栈顶index，注意 栈顶元素表示的是：上一个没有匹配成功的右括号的index
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
