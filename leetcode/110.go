package leetcode

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// 左右子树的高度差的平方小于1  而且左右子树分别都是平衡的
	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

/*
	104 二叉树最大深度
*/
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left)+1, height(root.Right)+1)
}

/*
	较大值
*/
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*
	取绝对值
*/
func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
