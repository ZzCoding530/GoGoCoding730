package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	var rootIndex int
	for k, v := range inorder {
		if v == preorder[0] {
			rootIndex = k
			break
		}
	}

	//这里直接用构造函数 很巧妙  具体的两个数组的切分需要自己画图加深记忆
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex]),
		Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])}
}
