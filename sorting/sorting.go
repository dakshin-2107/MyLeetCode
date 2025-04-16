package sorting

import "fmt"

func Sort() {

	//sortMethod := QuickSort
	//sortMethod := InsertionSort
	//sortMethod := SelectionSort
	sortMethod := MergeSort
	sortingInput := [][]int{
		{234, 345, 123, 456, 789, 69, 672, 985, 4098, 1234},
		{2, 3, 4, 4, 4, 5, 5, 1, 1, 1, 2, 2, 2},
	}

	for _, arr := range sortingInput {
		SortDelegate(arr, sortMethod)
	}
}

func SortDelegate(arr []int, sortMethod func([]int)) []int {
	fmt.Print("Input array : ", arr, "\n")
	sortMethod(arr)
	fmt.Print("Output array : ", arr, "\n\n")
	return arr
}

func PrintArray(arr []int) {
	for index, element := range arr {
		print(element)
		if index != len(arr)-1 {
			print(",")
		}
	}
}
