package leetcode

func sortArray(nums []int) []int {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(arr []int, low, high int) {
	if low >= high {
		return
	}

	index := partition(arr, low, high)
	quicksort(arr, low, index-1)
	quicksort(arr, index+1, high)
}

func partition(arr []int, low, high int) int {
	bound := arr[low]

	var i, j = low, high

	for i < j {
		for arr[j] >= bound && i < j {
			j--
		}

		for arr[i] <= bound && i < j {
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[i], arr[low] = arr[low], arr[j]
	return i
}
