package leetcode

/*
	和图相关，入度出度
*/
func findJudge(N int, trust [][]int) int {
	du := make([][]int, N)
	for i := range du {
		du[i] = []int{0, 0}
	}
	for _, v := range trust {
		x := v[0] - 1
		y := v[1] - 1
		du[x][0]++ // 出度++
		du[y][1]++ // 入度++
	}
	for i, v := range du {
		if v[0] == 0 && v[1] == N-1 {
			return i + 1
		}
	}
	return -1
}
