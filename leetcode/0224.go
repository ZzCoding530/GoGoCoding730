package leetcode

func calculate(s string) (ans int) {
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; { //遍历每一个字符
		switch s[i] {
		case ' ': //空字符往后走
			i++
		case '+': //加号的话
			sign = ops[len(ops)-1] //sign = 1
			i++
		case '-': //负号的话
			sign = -ops[len(ops)-1] //sign = -1
			i++
		case '(': //遇到左括号
			ops = append(ops, sign) //把当前sign放进去
			i++
		case ')': //遇到右括号
			ops = ops[:len(ops)-1] //
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return
}
