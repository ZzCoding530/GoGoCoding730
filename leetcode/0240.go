package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	col := 0

	for row >= 0 && col < len(matrix[0]) {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			row--
		} else if target > matrix[row][col] {
			col++
		}

	}

	return false
}
