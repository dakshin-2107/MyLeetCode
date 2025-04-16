package sorting

func Merge(lArr []int, rArr []int) []int {

	resArr := make([]int, 0)
	l := 0
	r := 0
	for {

		if r >= len(rArr) || l >= len(lArr) {
			break
		}

		if rArr[r] <= lArr[l] {
			resArr = append(resArr, rArr[r])
			r++
			continue
		}

		resArr = append(resArr, lArr[l])
		l++
	}

	for ; r < len(rArr); r++ {
		resArr = append(resArr, rArr[r])
	}

	for ; l < len(lArr); l++ {
		resArr = append(resArr, lArr[l])
	}

	return resArr
}

func DoMergeSort(arr []int, l, r int) []int {
	if l >= r {
		return arr[l : l+1]
	}

	mid := l + (r-l)/2

	leftArr := DoMergeSort(arr, l, mid)
	rightArr := DoMergeSort(arr, mid+1, r)
	return Merge(leftArr, rightArr)

}

func MergeSort(arr []int) {
	resArr := DoMergeSort(arr, 0, len(arr)-1)
	for index, _ := range arr {
		arr[index] = resArr[index]
	}
}
