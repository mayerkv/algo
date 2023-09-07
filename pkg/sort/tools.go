package sort

import "math/rand"

const seed = 12345

func Random(n int) []int {
	src := rand.NewSource(seed)
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

func Digits(n int) []int {
	src := rand.NewSource(seed)
	rnd := rand.New(src)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rnd.Intn(10)
	}
	return arr
}
