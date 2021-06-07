package leetcode

/*
	迭代法，其实指的是层序遍历的迭代版本， 如果是搜索二叉树，层序遍历应该是单调递增的
*/
func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}

	stack := []*TreeNode{root}
	result := []int{}

	for len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(result) > 1 {
			if isOrder(result) == false {
				return false
			}
		}

		result = append(result, root.Val)

		root = root.Right
	}

	return true

}

func isOrder(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}

	return true

}
