package finalShoot

import "strings"

func removeKdigits(num string, k int) string {
	stack := []byte{}    //新建一个队列
	for i := range num { //遍历nums
		digit := num[i]

		//循环条件 还没删够K个数  单调栈里面还有数  当前数字<栈顶元素
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit) // 入栈
	}
	stack = stack[:len(stack)-k] // 末尾的数字删除
	s := string(stack)
	s = strings.TrimLeft(s, "0") // 处理“00000xxx”这样的情况
	if s == "" {
		return "0"
	}
	return s
}
