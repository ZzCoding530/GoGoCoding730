package leetcode

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}

	sum := preSum*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return sum
	}

	return dfs(root.Left, sum) + dfs(root.Right, sum)

}
