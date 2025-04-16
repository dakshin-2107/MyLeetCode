package sorting

func SelectionSort(arr []int) {

	for j := 0; j < len(arr)-1; j++ {
		iterationMinIndex := -1
		iterationMin := 999999999

		for i := j; i < len(arr); i++ {
			if arr[i] < iterationMin {
				iterationMin = arr[i]
				iterationMinIndex = i
			}
		}

		arr[j], arr[iterationMinIndex] = arr[iterationMinIndex], arr[j]
	}
}
