package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Right
	root.Right = root.Left
	root.Left = temp

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
