package leetcode

var result int //final result to be return

func diameterOfBinaryTree(root *TreeNode) int {
	result = 0
	dfs(root)
	return result
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	result = max(result, left+right)

	return max(left, right) + 1

}
