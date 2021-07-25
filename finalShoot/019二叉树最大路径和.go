package finalShoot

import "math"

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32

	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := max(0, dfs(root.Left))
		r := max(0, dfs(root.Right))

		ans = max(ans, l+r+root.Val)

		return max(l, r) + root.Val
	}

	dfs(root)

	return ans
}
