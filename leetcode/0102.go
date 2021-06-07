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
		var nextLevel []*TreeNode // 存放下一层的节点
		var curLevelVal []int     // 存放当前层的节点值

		for _, node := range curLevel {
			curLevelVal = append(curLevelVal, node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		result = append(result, curLevelVal) //更新结果

		curLevel = nextLevel //当前层存的值更新给curLevel
	}

	return result
}
