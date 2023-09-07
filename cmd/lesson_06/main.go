package main

import (
	"fmt"
	"time"

	"github.com/mayerkv/algo/pkg/sort"
)

func main() {
	cases := []struct {
		name   string
		sorter sort.Sorter[int]
	}{
		{"Bubble", sort.BubbleSort[int]},
		{"Insertion", sort.Insertion[int]},
		{"InsertionShift", sort.InsertionShift[int]},
		{"InsertionBinary", sort.InsertionBinary[int]},
		{"Shell", sort.Shell[int]},
	}
	sizes := []int{100, 1_000, 10_000, 100_000}

	for _, size := range sizes {
		arr := sort.Random(size)
		for _, c := range cases {
			clone := make([]int, size)
			copy(clone, arr)
			now := time.Now()
			c.sorter(clone)
			fmt.Printf("%s(%d): %s\n", c.name, size, time.Since(now))
		}
		fmt.Println()
	}
}
