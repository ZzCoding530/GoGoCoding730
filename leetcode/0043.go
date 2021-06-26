package leetcode

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// num1[i] * num[j] 值必定在resArr[i+j] resArr[i+j+1]上，i+j+1存储低位
	resArr := make([]byte, len(num1)+len(num2))

	//i是num2的指针，从尾部开始指
	for i := len(num2) - 1; i >= 0; i-- {
		n2 := num2[i] - '0'
		//j是num1的指针，从尾部开始指
		for j := len(num1) - 1; j >= 0; j-- {
			n1 := num1[j] - '0'

			//index = i+j+1存低位， index = i+j存进的位数
			sum := n2*n1 + resArr[i+j+1]
			resArr[i+j+1] = sum % 10
			resArr[i+j] += sum / 10
		}
	}

	//去开头的0
	j := -1
	for i := 0; i < len(num1)+len(num2); i++ {

		// 如果字符不等于0   且	j==-1
		if resArr[i] != 0 && j == -1 {
			j = i
		}
		resArr[i] += '0'
	}

	return string(resArr[j:])
}
