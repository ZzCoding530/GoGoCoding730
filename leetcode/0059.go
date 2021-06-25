package leetcode

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	num := 1                                   //起始值为1
	left, right, top, bottom := 0, n-1, 0, n-1 // 四条边的边界

	for num <= n*n {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++

		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}

	return res
}
