package leetcode

import "../structures"

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
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
		level = temp
	}

	return result

}
