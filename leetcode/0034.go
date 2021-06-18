package leetcode

func searchRange(nums []int, target int) []int {
	left := search(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := search(nums, target+1) - 1
	return []int{left, right}
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	// [l, r]
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
