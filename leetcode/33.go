package leetcode

/*
	第一类 2 3 4 5 6 7 1 这种，也就是 nums[start] <= nums[mid]。此例子中就是 2 <= 5。
	这种情况下，前半部分有序。因此如果 nums[start] <=target<nums[mid]，则在前半部分找，否则去后半部分找。
	第二类 6 7 1 2 3 4 5 这种，也就是 nums[start] > nums[mid]。此例子中就是 6 > 2。
	这种情况下，后半部分有序。因此如果 nums[mid] <target<=nums[end]，则在后半部分找，否则去前半部分找。

	!!!!!说白了，就是无限分割永远都在有序的序列里面找!!!!!
*/
func search(nums []int, target int) int {
	low, high, mid := 0, len(nums)-1, 0

	for low <= high {
		mid = low + (high-low)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[low] { //这句主要是判断mid这个数是在左段还是右段
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
