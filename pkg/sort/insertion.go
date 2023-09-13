package sort

func Insertion[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0 && arr[i] > arr[i+1]; i-- {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}

func InsertionShift[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
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
	if n <= 1 {
		return
	}
	var i int
	for j := 1; j < n; j++ {
		k := arr[j]
		for i = j - 1; i >= 0 && arr[i] > k; i-- {
			arr[i+1] = arr[i]
		}
		arr[i+1] = k
	}
}
