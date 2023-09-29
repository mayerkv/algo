package sort

func CountSort[T Int](arr []T) {
	n := len(arr)

	var max T
	for i := 0; i < n; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	countArray := make([]int, max+1)
	for _, v := range arr {
		countArray[v] += 1
	}

	for i := range countArray {
		if i == 0 {
			continue
		}
		countArray[i] += countArray[i-1]
	}

	outputArray := make([]T, n)
	for i := n - 1; i >= 0; i-- {
		outputArray[countArray[arr[i]]-1] = arr[i]
		countArray[arr[i]]--
	}

	copy(arr, outputArray)
}
