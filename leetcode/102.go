package leetcode

import "../structures"

type TreeNode = structures.TreeNode

func levelOrder(root *TreeNode) [][]int {
	var result [][]int // 结果数组，二维的

	if root == nil {
		return result
	}

	curLevel := []*TreeNode{root} // 存放当前层的节点

	for len(curLevel) > 0 { // 直到当前层没有节点
		nextLevel := []*TreeNode{} // 存放下一层的节点
		curLevelVals := []int{}    // 存放当前层的节点值

		for _, node := range curLevel {
			curLevelVals = append(curLevelVals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		result = append(result, curLevelVals)

		curLevel = nextLevel
	}

	return result
}
