package leetcode

/*
	从后往前指针走，从后往前确定数组每一个格子的数
	https://leetcode-cn.com/problems/merge-sorted-array/solution/88-by-ikaruga/
	这个题解非常明白
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	index := len(nums1) - 1
	m--
	n--

	for n >= 0 {
		for m >= 0 && nums1[m] > nums2[n] {
			nums1[index], nums1[m] = nums1[m], nums1[index]
			index--
			m--
		}

		nums1[index], nums2[n] = nums2[n], nums1[index]
		index--
		n--
	}

}
