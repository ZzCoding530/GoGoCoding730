package leetcode

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		tempRoot := queue[0]
		queue = queue[1:]

		tempRoot.Left, tempRoot.Right = tempRoot.Right, tempRoot.Left

		if tempRoot.Left != nil {
			queue = append(queue, tempRoot.Left)
		}

		if tempRoot.Right != nil {
			queue = append(queue, tempRoot.Right)
		}
	}

	return root
}
