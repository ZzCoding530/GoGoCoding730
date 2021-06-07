package leetcode

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k)
}

func quickSort(nums []int, low, high, index int) int {
	pIndex := partition(nums, low, high)
	if pIndex == len(nums)-index {
		return nums[pIndex]
	}

	if pIndex < len(nums)-index {
		return quickSort(nums, pIndex+1, high, index)
	} else {
		return quickSort(nums, low, pIndex-1, index)
	}

}

func partition(nums []int, low, high int) int {
	bound := nums[low]
	l, r := low, high

	for l < r {
		for l < r && nums[r] >= bound {
			r--
		}

		for l < r && nums[l] <= bound {
			l++
		}

		nums[l], nums[r] = nums[r], nums[l]

	}

	nums[low], nums[l] = nums[l], nums[low]

	return l
}
