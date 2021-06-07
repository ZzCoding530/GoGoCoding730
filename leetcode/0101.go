package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root.Left, root.Right)
}

func dfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return dfs(left.Left, right.Right) && dfs(right.Left, left.Right)
}
