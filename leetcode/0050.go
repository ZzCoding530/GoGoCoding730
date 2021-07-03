package leetcode

/*
	利用快速幂求值，注意处理n正负时候的区别
*/
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

/*
	快速幂取法
*/
func quickMul(x float64, n int) float64 {
	if n == 0 { //0次幂是1
		return 1
	}

	y := quickMul(x, n/2) //先求一半的幂（幂次数自动向下取整）
	if n%2 == 0 {         //如果原先的幂是偶数，那没事直接 n次幂 = (n/2)次幂 * (n/2)次幂
		return y * y
	}
	return y * y * x //如果原先是奇数次幂，那么需要多乘一次幂 即自己
}
