package leetcode

import "sort"

func nextPermutation(nums []int) {
	l := len(nums)

	for i := l - 1; l > 0; l-- {
		if nums[i] > nums[i-1] {
			sort.Ints(nums[i:])
			for j := i; j < l; j++ {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					return
				}
			}
		}
	}

	sort.Ints(nums)
	return
}
