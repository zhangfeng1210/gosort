package msort

//Bubble sort
func SortValues(values []int) {
	flag := true
	for i := 0; i < len(values); i++ {
		flag = true
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				flag = false
			}
		}
		if flag == true {
			break
		}
	}
}

//quick sort
func QuickSort(values []int, start, end int) {
	if start > end {
		return
	}

	left, p := start, start
	right := end
	temp := values[p]

	for right > left {
		for values[right] > temp {
			right--
		}
		values[right], values[p] = values[p], values[right]

		p = right
		for values[left] < temp {
			left++
		}
		values[left], values[p] = values[p], values[left]

		p = left
	}

	QuickSort(values, start, left-1)
	QuickSort(values, right+1, end)

}
