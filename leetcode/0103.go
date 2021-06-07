package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}

	level := []*TreeNode{root}

	for i := 0; len(level) > 0; i++ {
		result = append(result, []int{})
		var temp []*TreeNode

		for j := 0; j < len(level); j++ {
			node := level[j]
			result[i] = append(result[i], node.Val)
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}

		if i%2 == 1 {
			for a, b := 0, len(temp)-1; a < b; a, b = a+1, b-1 {
				temp[a], temp[b] = temp[b], temp[a]
			}
		}
		level = temp
	}

	return result
}
