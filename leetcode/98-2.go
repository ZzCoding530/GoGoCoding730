package leetcode

var list *[]int = new([]int)

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}
	list = inorder(root)

	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}

	return true

}

func inorder(root *TreeNode) []int {
	if root != nil {
		inorder(root.Left)
		list = append(list, root.Val)
		inorder(root.Right)
	}
	return list
}
