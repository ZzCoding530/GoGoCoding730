package leetcode

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int //存最终结果的二维int切片
	var tempPath []int //存每个路径的一维int切片

	return dfs(root, tempPath, result, targetSum)
}

//         根节点			   当前的路径	   最后返回的结果   目标sum
func dfs(root *TreeNode, tempPath []int, result [][]int, sum int) [][]int {
	if root == nil {
		return result
	}

	if root.Left == nil && root.Right == nil { //假如到了最后的子节点
		if sum == 0 { //到最下面以后如果需要的sum已经是0 了 那就是找到了  当前的tempPath就是一条了
			tempSlice := make([]int, len(tempPath)) //新建一个切片
			copy(tempSlice[0:], tempPath[0:])       //把当前路径复制过去
			result = append(result, tempSlice)      //把新复制好的放到最终结果里

		}

		return result
	}

	if root.Left != nil { //如果没走到最下面，那按照前序遍历，先往左
		result = dfs(root.Left, tempPath, result, sum)
	}

	if root.Right != nil { //再往右
		result = dfs(root.Right, tempPath, result, sum)
	}

	return result
}
