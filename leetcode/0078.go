package leetcode

func subsets(nums []int) (ans [][]int) {
	set := []int{} //暂时的子集
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) { //全都遍历了 返回吧
			ans = append(ans, append([]int(nil), set...)) //append([]int(nil), set...)这个写法相当于是新建一个空切片然后把set当前的值复制进去，避免下一轮回溯set变了
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
