package leetcode

/*
	思路，层序遍历 遍历到nil以后就不应该再遍历到不是nil的节点了
*/
func isCompleteTree(root *TreeNode) bool {
	levelNode := []*TreeNode{root}
	lastIsNil := false

	for len(levelNode) > 0 {
		node := levelNode[0]
		levelNode = levelNode[1:]

		if node == nil {
			lastIsNil = true
		} else {
			if lastIsNil {
				return false
			}

			levelNode = append(levelNode, node.Left)
			levelNode = append(levelNode, node.Right)
		}

	}

	return true
}
