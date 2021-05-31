package leetcode

type item struct { //item结构体，idx是层序遍历的序号，还有节点本身
	idx int
	*TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans, que := 1, []item{{0, root}} //初始化一下结果，和层序遍历用到的队列
	for len(que) > 0 {
		if l := que[len(que)-1].idx - que[0].idx + 1; l > ans { //如果宽度比ans大了，更新一下ans
			ans = l
		}

		var tmp []item //新建一个放item的切片
		for _, q := range que {
			if q.Left != nil {
				tmp = append(tmp, item{q.idx * 2, q.Left})
			}
			if q.Right != nil {
				tmp = append(tmp, item{q.idx*2 + 1, q.Right})
			}
		}
		que = tmp
	}
	return ans
}

/*
	其实总思路就是：在层序的基础上，把每个节点在层序遍历中的序号带上了，每一层都计算更新一下
*/
