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

	for size := 100; size <= 1_000_000; size *= 10 {
		arr := sort.Random(size)
		timeout := getTimeoutForSize(size)
		for _, c := range cases {
			d, err := withTimeout(c.sorter, arr, timeout)
			if err != nil {
				fmt.Printf("%s(%d): %s\n", c.name, size, err)
			} else {
				fmt.Printf("%s(%d): %s\n", c.name, size, d)
			}
		}
		fmt.Println()
	}
}

func getTimeoutForSize(size int) time.Duration {
	timeout := 15 * time.Second
	if size > 100_000 {
		timeout = 2 * time.Minute
	}
	return timeout
}

func withTimeout(sorter sort.Sorter[int], arr []int, timeout time.Duration) (time.Duration, error) {
	clone := copyArr(arr)
	ch := handle(sorter, clone)
	timer := time.NewTimer(timeout)

	select {
	case d := <-ch:
		return d, nil
	case <-timer.C:
		return 0, fmt.Errorf("timeout")
	}
}

func handle(fn sort.Sorter[int], arr []int) chan time.Duration {
	ch := make(chan time.Duration, 1)
	go func() {
		defer close(ch)
		now := time.Now()
		fn(arr)
		ch <- time.Since(now)
	}()
	return ch
}

func copyArr(arr []int) []int {
	clone := make([]int, len(arr))
	copy(clone, arr)
	return clone
}
