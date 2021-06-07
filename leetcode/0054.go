package leetcode

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var res []int //存结果的

	top := 0 //上下左右的边界
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top < bottom && left < right {
		for i := left; i < right; i++ { //先从左上到右上呗
			res = append(res, matrix[top][i])
		}
		for i := top; i < bottom; i++ { //然后从右上到右下
			res = append(res, matrix[i][right])
		}
		for i := right; i > left; i-- { //然后从右下到左下
			res = append(res, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- { //然后从左下到左上
			res = append(res, matrix[i][left])
		}
		//往里缩一圈
		right--
		top++
		bottom--
		left++
	}
	if top == bottom { //top == bottom 的时候大概是，只剩中间的一行了，不构成矩形，就是一行，这一行都加进去就行
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	} else if left == right { //left == right 的时候应该是，只剩下中间的一列可，不构成矩形就是一列，这一列就加进去
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][left])
		}
	}
	return res
}
