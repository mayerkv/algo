package sort

func BubbleSort[T Ordered](arr []T) {
	n := len(arr)
	if n <= 1 {
		return
	}
	for j := n - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
}
