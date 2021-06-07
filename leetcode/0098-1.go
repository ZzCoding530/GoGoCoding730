package leetcode

import "math"

/*
	递归
*/
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

/*
	lower upper 上下界
*/
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	//当前节点是左子树的上界									是右子树的下界
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}
