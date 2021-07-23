package finalShoot

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := []*Node{}          //这一层的nodes
	queue = append(queue, root) //首先加入root节点

	for len(queue) != 0 {
		thisLevelSize := len(queue) //这一层有多少节点
		thisLevelVals := []int{}
		for i := 0; i < thisLevelSize; i++ {
			tempNode := queue[0] //从queue里取出头节点
			queue = queue[1:]
			thisLevelVals = append(thisLevelVals, tempNode.Val)

			childrenNodes := tempNode.Children
			for j := 0; j < len(childrenNodes); j++ {
				if childrenNodes[j] != nil {
					queue = append(queue, childrenNodes[j])
				}
			}

		}
		ans = append(ans, thisLevelVals)
	}

	return ans

}
