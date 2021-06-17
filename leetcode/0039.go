package leetcode

func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) { //走到最后一个数了 return
			return
		}
		if target == 0 { //已经和为target了 return当前的组合
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		if target-candidates[idx] >= 0 { // 选择当前数
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}

		dfs(target, idx+1) //跳过当前数
	}
	dfs(target, 0)
	return //返回值是有名字的就不用在函数内返回了
}
