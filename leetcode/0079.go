package leetcode

func exist(board [][]byte, word string) bool {
	var dfs func(int, int, int) bool

	/*
		这个函数表示  从(x,y)出发，下一步能不能碰到word[index]
	*/
	dfs = func(y, x, index int) bool {
		if index == len(word) {
			return true
		}
		if y < 0 || x < 0 || y == len(board) || x == len(board[y]) { //如果出界了 false
			return false
		}
		if board[y][x] != word[index] { //如果当前位置的字母不是想要的 false
			return false
		}

		temp := board[y][x] //暂存字符用于后面回溯恢复
		board[y][x] = '#'   //先把字符改为#，表示这个位置已经来过了

		exist := dfs(y, x+1, index+1) || dfs(y, x-1, index+1) || dfs(y+1, x, index+1) || dfs(y-1, x, index+1) //然后看存不存在，上下左右又一个存在就行了

		board[y][x] = temp //这个位置恢复一下
		return exist
	}

	//然后就遍历每个格子呗
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if dfs(y, x, 0) {
				return true
			}
		}
	}
	return false
}
