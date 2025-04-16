package sorting

// start from left and swap elements that are less than the pivot to the left
func lomutoPartition(arr []int, low, high int) int {

	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	i++
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

// traverse from both sides until order is not valid and then swap the elements
func hoarePartition(arr []int, low, high int) int {

	pivot := arr[low]
	pivotloc := low
	low++

	for {

		for arr[high] > pivot {
			high--
		}

		for low <= high && arr[low] < pivot {
			low++
		}

		if low >= high {
			arr[high], arr[pivotloc] = arr[pivotloc], arr[high]
			return high
		}

		arr[low], arr[high] = arr[high], arr[low]
	}
}

func DoQuickSort(arr []int, low, high int) {

	if low < high {
		pi := lomutoPartition(arr, low, high)
		//pi := hoarePartition(arr, low, high)
		DoQuickSort(arr, low, pi-1)
		DoQuickSort(arr, pi+1, high)
	}
}

func QuickSort(arr []int) {
	DoQuickSort(arr, 0, len(arr)-1)
}
