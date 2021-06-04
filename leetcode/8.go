package leetcode

import "math"

func myAtoi(s string) int {
	index := 0 //字符指针
	sign := 1  // 正负号
	ans := 0   //结果
	l := len(s)

	for index < l && s[index] == ' ' { //去掉前面多余的空格
		index++
	}

	if index < l { //去掉多余空格以后 下一位应该就是正负号了
		if s[index] == '+' {
			sign = 1
			index++
		} else if s[index] == '-' {
			sign = -1
			index++
		}
	}

	for index < l && s[index] <= '9' && s[index] >= '0' { //取下一位,如果是代表数字的字符
		ans = ans*10 + int(s[index]-'0') //通过给旧的ans*10来使其左移，然后加上目前这一位字符代表的数字

		if ans*sign > math.MaxInt32 { //别忘了判断当前是不是溢出了
			return math.MaxInt32
		} else if ans*sign < math.MinInt32 {
			return math.MinInt32
		}

		index++
	}

	return ans * sign
}
