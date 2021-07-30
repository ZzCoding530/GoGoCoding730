package finalShoot

/*
	快速排序最重要的一点是partition过程
*/
func sortArray(nums []int) []int {

}

func quickSort(nums []int, low, high int) {
	if low >= high {
		return
	}

	index := partition(nums, low, high)
	quickSort(nums, low, index-1)
	quickSort(nums, index+1, high)
}

func partition(nums []int, low, high int) int {
	bound := low

	var i, j = low, high

	for i < j {
		for nums[j] >= bound && i < j {
			j--
		}

		for nums[i] <= bound && i < j {
			i++
		}

		nums[i], nums[j] = nums[j], nums[i]

	}

	nums[i], nums[low] = nums[low], nums[i]
	return i
}
