package finalShoot

func search(nums []int, target int) int {
	l, r, mid := 0, len(nums), len(nums)/2
	for r >= l {
		if target == nums[mid] {
			return nums[mid]
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
		mid = (l + r) / 2
	}
	return -1
}
