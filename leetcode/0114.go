package leetcode

/*
	将左右子树分别递归展开，将原左子树变为节点的右子树，再将原右子树变为当前右子树最右节点的右子树
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	//此时左右子树已经被分别拉成一条线了
	r := root.Right
	root.Right, root.Left = root.Left, nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = r
}
