package sort

import (
	"fmt"
	"testing"
)

func BenchmarkSort(b *testing.B) {
	inputSize := []int{100, 1_000, 10_000, 100_000}
	generators := []struct {
		name string
		fn   func(int) []int
	}{
		{"Random", Random},
		{"Digits", Digits},
		{"Sorted", Sorted},
		{"Reversed", Reversed},
	}
	cases := []struct {
		name string
		sort Sorter[int]
	}{
		{
			name: "Bubble",
			sort: BubbleSort[int],
		},
		{
			name: "Insertion",
			sort: Insertion[int],
		},
		{
			name: "InsertionShift",
			sort: InsertionShift[int],
		},
		{
			name: "InsertionBinary",
			sort: InsertionBinary[int],
		},
		{
			name: "Shell",
			sort: Shell[int],
		},
		{
			name: "SelectionSort",
			sort: SelectionSort[int],
		},
		{
			name: "HeapSort",
			sort: HeapSort[int],
		},
		{
			name: "MergeSort",
			sort: MergeSort[int],
		},
		{
			name: "QuickSort",
			sort: QuickSort[int],
		},
	}
	for _, g := range generators {
		b.Run(
			g.name, func(b *testing.B) {
				for _, size := range inputSize {
					b.Run(
						fmt.Sprint(size), func(b *testing.B) {
							arr := g.fn(size)
							b.ResetTimer()
							for _, c := range cases {
								b.Run(
									c.name, func(b *testing.B) {
										for i := 0; i < b.N; i++ {
											clone := make([]int, size)
											copy(clone, arr)
											c.sort(clone)
										}
									},
								)
							}
						},
					)
				}
			},
		)

	}
}
