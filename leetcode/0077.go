package leetcode

var res [][]int

func combine(n int, k int) [][]int {
	//再初始化一下，不然多个测试会互相影响
	res = [][]int{}

	if n <= 0 || k <= 0 || k > n {
		return res
	}
	back(n, k, 1, []int{})

	return res
}

func back(n int, k int, startIndex int, temp []int) {

	//这地方有个剪枝，说白了判断剩下的数都放进去够不够长度k，不够就没必要继续了
	if len(temp)+n-startIndex+1 < k {
		return
	}

	//判断是不是已经集齐k个组合了，这里注意要新建一个切片副本，不然切片会被下一次操作改变
	if len(temp) == k {
		tempAns := make([]int, k)
		copy(tempAns, temp)
		res = append(res, tempAns)
	}

	for i := startIndex; i <= n; i++ {
		temp := append(temp, i)
		back(n, k, i+1, temp)
		temp = temp[:len(temp)-1]
	}
}
