package leetcode

func kthLargest(root *TreeNode, k int) int {
	var nums []int
	nums = getNums(&nums, root)
	return nums[k-1]
}

//逆着中序遍历  右中左
func getNums(nums *[]int, r *TreeNode) []int {
	if r.Right != nil {
		getNums(nums, r.Right)
	}
	if r != nil {
		*nums = append(*nums, r.Val)
	}
	if r.Left != nil {
		getNums(nums, r.Left)
	}
	return *nums
}
