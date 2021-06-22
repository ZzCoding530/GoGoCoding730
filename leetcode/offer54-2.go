package leetcode

/*
	这种不需要数组保存值，要熟练应用函数变量的写法
*/
func kthLargest(root *TreeNode, k int) int {
	res := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil && k > 0 {
			dfs(root.Right)
			k--
			if k == 0 {
				res = root.Val
			}
			dfs(root.Left)
		}
	}
	dfs(root)
	return res
}
