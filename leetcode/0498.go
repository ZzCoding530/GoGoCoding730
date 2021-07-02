package leetcode

func findDiagonalOrder(mat [][]int) []int {
	m := len(mat)
	n := len(mat[0])

	result := make([]int, m*n)
	i := 0
	j := 0
	direction := 1

	for curIndex := 0; curIndex < len(result); curIndex++ {
		result[curIndex] = mat[i][j]
		if j == n-1 && direction == 1 {
			direction = -1
			i = i + 1
			continue
		}
		if i == m-1 && direction == -1 {
			direction = 1
			j = j + 1
			continue
		}
		if i == 0 && direction == 1 {
			direction = -1
			j = j + 1
			continue
		}
		if j == 0 && direction == -1 {
			direction = 1
			i = i + 1
			continue
		}
		i = i - direction
		j = j + direction
	}

	return result
}
