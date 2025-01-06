package main

import "fmt"

func MyFunc(nums []int, target int) []int {

	var itemsMap = make(map[int]int)

	for index, item := range nums {

		diff := target - item
		fmt.Println(diff)
		if secondIndex, isPresent := itemsMap[diff]; isPresent {
			return []int{secondIndex, index}
		}
		itemsMap[item] = index
		fmt.Println(itemsMap)
	}

	return []int{}
}

func main() {

	nums := [...]int{3, 2, 4}
	target := 6

	fmt.Println(MyFunc(nums[:], target))
}
