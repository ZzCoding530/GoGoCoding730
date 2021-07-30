package finalShoot

/*
	很具小巧思。新建数组 cell，用于保存最长上升子序列。

	对原序列进行遍历，将每位元素二分插入 cell 中。

	如果 cell 中元素都比它小，将它插到最后
	否则，用它覆盖掉比它大的元素中最小的那个。
	总之，思想就是让 cell 中存储比较小的元素。这样，cell 未必是真实的最长上升子序列，但长度是对的。

*/
func lengthOfLIS(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	cell := []int{}

	cell = append(cell, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > cell[len(cell)-1] {
			cell = append(cell, nums[i])
			continue
		}

		l, r := 0, len(cell)-1
		for l < r {
			mid := l + (r-l)/2
			if cell[mid] < nums[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}

		cell[l] = nums[i]

	}

	return len(cell)

}
