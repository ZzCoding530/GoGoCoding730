package leetcode

func rand10() int {
	for {
		r1 := rand7()
		r2 := rand7()
		num := r1 + (r2-1)*7
		if num <= 40 {
			return num%10 + 1
		}
	}
}

/*
	一些思考 如果用04的生成09的那么
	0-4的随机生成可以看成是0-3的随机生成，就不想要的直接舍弃，然后03*3就成了09
	反正就故意舍弃，然后乘以常数和加减常数即可
*/
