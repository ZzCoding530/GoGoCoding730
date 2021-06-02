package leetcode

func search(nums []int, target int) int {
	pivot, left, right := 0, 0, len(nums)-1

	for left < right {
		pivot = left + (right-left)/2

		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}

	}

	return -1
}
