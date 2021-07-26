package finalShoot

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	i := 0                        //i 要定义在外面，因为for循环后还要用
	for ; i < len(inorder); i++ { //从inorder数组中找到preorder[0]的位置
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root

}
