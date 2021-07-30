package finalShoot

/*
	二分法，其实跟直接找第一个nums[i]<nums[i-1]速度没差。。。都是100%
*/
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}
