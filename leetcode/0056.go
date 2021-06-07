package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { //按照区间的左边界排序，升序排列
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int     //存结果的
	std := intervals[0] //被拿来比较的第一个区间，作为 基准区间  注意这是排序后排在第一个的区间

	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]  //当前的区间
		if std[1] < cur[0] { // if拿来对比的区间的右边界 < 当前区间的左边界  那就说明两个区间没有一点重合
			res = append(res, std) //那就说明当前的区间是一个独立的区间
			std = cur              // 因为是排序过的，所以既然这个区间都没重合部分，那就可以把这个区间当作基准，比对下一个区间了
		} else { // 有重合
			std[1] = max(std[1], cur[1]) //如果有重合，即【1，3】【2，6】这种，因为求的是并集，那就把右边界更新成大的，不能直接赋值，要考虑【1，3】【2，3】这种情况
		}
	}
	res = append(res, std) //把更新完右边界的这个基准区间放到答案里
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
