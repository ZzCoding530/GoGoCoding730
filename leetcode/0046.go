package leetcode

func permute(nums []int) [][]int {
	used := map[int]bool{} //默认是false
	var ans [][]int

	var dfs func(path []int)
	dfs = func(path []int) {

		if len(path) == len(nums) { //判断是不是到最后了，path里面的数字已经全了
			temp := make([]int, len(path))
			copy(temp, path) //全了那就复制一份放到结果里
			ans = append(ans, temp)
			return
		}

		for _, n := range nums { //正常循环逻辑在这里
			if used[n] { //如果这个数字被用过了，那就跳过去看下一位
				continue
			}
			path = append(path, n)    //没用过的数字直接加进path里
			used[n] = true            //因为用过了所以要置成true
			dfs(path)                 //path传进去进行下一步递归
			path = path[:len(path)-1] //还要进行不传进去的操作，拿出来
			used[n] = false           //置成false，假装直接跳过了这一个数进行下一个for循环，但其实没跳过，而是刚才新开了一根分枝正在自动往下dfs
		}
	}

	dfs([]int{})
	return ans

}
