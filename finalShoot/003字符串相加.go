package finalShoot

func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	add := 0
	l := max(len1, len2)

	ansStr := make([]byte, l+1)

	for i := l - 1; i >= 0; i-- { //从后往前遍历
		i1 := i - (l - len1) //这两句是为了都让从最后开始遍历
		i2 := i - (l - len2) //l-len 是补齐用的
		a := 0
		b := 0
		if i1 >= 0 { //如果短的那个到头了就是默认值 0，另一个数自己加
			a = int(num1[i1] - '0')
		}
		if i2 >= 0 {
			b = int(num2[i2] - '0')
		}
		ansStr[i+1] = byte(((a + b + add) % 10) + '0') //因为ansStr预留了一位最前面防止进位用的，所以保存时候往后挪一位
		add = (a + b + add) / 10                       //算进位

	}
	if add == 0 { //如果最后没进位，把预留的截掉返回
		return string(ansStr[1:])
	}

	ansStr[0] = '1' //最后还有进位，加上，返回
	return string(ansStr)

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
