package leetcode

/*
	迭代法 就相当于层序遍历那种 一行一行翻，从上到下从左到右一个一个做root翻
*/
func invertTree(root *TreeNode) *TreeNode {
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
