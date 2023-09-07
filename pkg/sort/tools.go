package sort

import "math/rand"

func Random(n int) []int {
	src := rand.NewSource(12345)
	rnd := rand.New(src)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rnd.Intn(n)
	}
	return arr
}

func Sorted(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func Reversed(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - 1 - i
	}
	return arr
}
