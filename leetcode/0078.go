package leetcode

func subsets(nums []int) (ans [][]int) {
	set := []int{} //暂时的子集
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) { //全都遍历了 返回吧
			temp := []int{}
			temp = append(temp, set...)
			ans = append(ans, temp)
			//ans = append(ans, append([]int(nil), set...))  这句可以代替上面三句
			return
		}

		set = append(set, nums[cur]) //首先把当前值放进set
		dfs(cur + 1)                 //然后下一个值
		set = set[:len(set)-1]       //然后把刚才放进去的拿出来
		dfs(cur + 1)                 //重复刚才的
	}
	dfs(0)
	return
}
