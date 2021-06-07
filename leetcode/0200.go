package leetcode

func numIslands(grid [][]byte) int {
	ans := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ { //总的来说，遍历每个
			if grid[row][col] == '1' {
				ans++                           //每碰到一个格子是'1' 就说明发现了一个岛屿
				disAppearIsland(grid, row, col) //为了不重复发现岛屿，我们需要把与刚发现的这个格子连着的格子都置为'2'
			}
		}
	}

	return ans
}

func disAppearIsland(grid [][]byte, row, col int) {
	if row >= len(grid) || row < 0 || col >= len(grid[0]) || col < 0 { //如果出界了就不用了 return
		return
	}

	if grid[row][col] != '1' { //本来就不是岛屿的话，也不必操作了 return
		return
	}

	grid[row][col] = '2' //这一步是真正的核心，把岛屿'1'置为非岛屿'2'的一步

	//递归这个格子连着的上下左右四个格子
	disAppearIsland(grid, row-1, col)
	disAppearIsland(grid, row+1, col)
	disAppearIsland(grid, row, col-1)
	disAppearIsland(grid, row, col+1)

}
