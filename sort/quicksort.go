package sort

func quickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, low, high int) {
	if low >= high {
		return
	}

	partitionIndex := partition(arr, low, high)
	sort(arr, low, partitionIndex-1)
	sort(arr, partitionIndex+1, high)

}

func partition(arr []int, low, high int) int {
	bound := arr[low]

	var i, j int = low, high

	for i < j {
		for arr[j] >= bound && i < j {
			j--
		}

		for arr[i] <= bound && i < j {
			i++
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[i], arr[low] = arr[low], arr[i]
	return i
}
