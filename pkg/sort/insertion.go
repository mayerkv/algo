package sort

func Insertion[T Ordered](arr []T) {
	n := len(arr)
	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0 && arr[i] > arr[i+1]; i-- {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}

func InsertionShift[T Ordered](arr []T) {
	n := len(arr)
	var i int
	for j := 1; j < n; j++ {
		k := arr[j]
		p := binarySearch(arr, k, 0, j-1)
		for i = j - 1; i >= p; i-- {
			arr[i+1] = arr[i]
		}
		arr[i+1] = k
	}
}

func InsertionBinary[T Ordered](arr []T) {
	n := len(arr)
	var i int
	for j := 1; j < n; j++ {
		k := arr[j]
		for i = j - 1; i >= 0 && arr[i] > k; i-- {
			arr[i+1] = arr[i]
		}
		arr[i+1] = k
	}
}

func binarySearch[T Ordered](arr []T, key T, low, high int) int {
	if high <= low {
		if key >= arr[low] {
			return low + 1
		}
		return low
	}
	mid := (low + high) / 2
	if key >= arr[mid] {
		return binarySearch(arr, key, mid+1, high)
	}
	return binarySearch(arr, key, low, mid-1)
}
