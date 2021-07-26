package finalShoot

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	var res []int //存结果的

	top := 0 //上下左右的边界
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for left < right && top < bottom {

		for i := left; i < right; i++ {
			res = append(res, matrix[top][i])
		}

		for i := top; i < bottom; i++ {
			res = append(res, matrix[i][right])
		}

		for i := right; i > left; i-- {
			res = append(res, matrix[bottom][i])
		}

		for i := bottom; i > top; i-- {
			res = append(res, matrix[i][left])
		}

		left++
		right--
		bottom--
		top++

	}

	if left == right {
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][left])
		}
	} else if top == bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
	}

	return res

}
