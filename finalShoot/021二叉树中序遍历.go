package finalShoot

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil { //先放左子树，到尽头
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		root = root.Right

	}

	return ans
}
