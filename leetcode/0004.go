package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var n int = len(nums1)
	var m int = len(nums2)
	left := (n + m + 1) / 2
	right := (n + m + 2) / 2
	return float64(getKth(nums1, 0, n-1, nums2, 0, m-1, left)+getKth(nums1, 0, n-1, nums2, 0, m-1, right)) * 0.5
}

func getKth(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	//让 len1 的长度小于 len2，这样就能保证如果有数组空了，一定是 len1，属于一个小trick方便后面判断
	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}

	if len1 == 0 { //nums1空的就直接从nums2里找
		return nums2[start2+k-1]
	}

	if k == 1 { //如果找第1小的，就是找最小值
		return min(nums1[start1], nums2[start2])
	}

	// 取 len1 和 k/2 小 是为了防止数组长度小于 k/2
	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}
