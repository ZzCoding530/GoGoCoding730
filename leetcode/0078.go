package leetcode

func subsets(nums []int) (ans [][]int) {
	set := []int{} //暂时的子集
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) { //全都遍历了 返回吧
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}
