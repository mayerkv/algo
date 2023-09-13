package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mayerkv/algo/pkg/sort"
	"github.com/mayerkv/algo/pkg/tester"
)

func main() {
	runTests()
}

func report() {
	cases := []struct {
		name   string
		sorter sort.Sorter[int]
	}{
		//{"Bubble", sort.BubbleSort[int]},
		//{"Insertion", sort.Insertion[int]},
		//{"InsertionShift", sort.InsertionShift[int]},
		//{"InsertionBinary", sort.InsertionBinary[int]},
		//{"Shell", sort.Shell[int]},
		{"SelectionSort", sort.SelectionSort[int]},
		{"HeapSort", sort.HeapSort[int]},
	}

	for size := 100; size <= 1_000_000; size *= 10 {
		arr := sort.Random(size)
		timeout := getTimeoutForSize(size)
		for _, c := range cases {
			d, err := withTimeout(c.sorter, copyArr(arr), timeout)
			if err != nil {
				fmt.Printf("%s(%d): %s\n", c.name, size, err)
			} else {
				fmt.Printf("%s(%d): %s\n", c.name, size, d)
			}
		}
		fmt.Println()
	}
}

func runTests() {
	tests := []*tester.Tester{
		//tester.NewTester(&sortTask{fn: sort.BubbleSort[int]}, "/Users/kmayer/Downloads/sorting-tests/0.random"),
		//tester.NewTester(&sortTask{fn: sort.Insertion[int]}, "/Users/kmayer/Downloads/sorting-tests/0.random"),
		//tester.NewTester(&sortTask{fn: sort.InsertionShift[int]}, "/Users/kmayer/Downloads/sorting-tests/0.random"),
		//tester.NewTester(&sortTask{fn: sort.InsertionBinary[int]}, "/Users/kmayer/Downloads/sorting-tests/0.random"),
		//tester.NewTester(&sortTask{fn: sort.Shell[int]}, "/Users/kmayer/Downloads/sorting-tests/0.random"),
		tester.NewTester(&sortTask{fn: sort.SelectionSort[int]}, "/Users/kmayer/Downloads/sorting-tests/0.random"),
		tester.NewTester(&sortTask{fn: sort.HeapSort[int]}, "/Users/kmayer/Downloads/sorting-tests/0.random"),

		//tester.NewTester(&sortTask{fn: sort.BubbleSort[int]}, "/Users/kmayer/Downloads/sorting-tests/1.digits"),
		//tester.NewTester(&sortTask{fn: sort.Insertion[int]}, "/Users/kmayer/Downloads/sorting-tests/1.digits"),
		//tester.NewTester(&sortTask{fn: sort.InsertionShift[int]}, "/Users/kmayer/Downloads/sorting-tests/1.digits"),
		//tester.NewTester(&sortTask{fn: sort.InsertionBinary[int]}, "/Users/kmayer/Downloads/sorting-tests/1.digits"),
		//tester.NewTester(&sortTask{fn: sort.Shell[int]}, "/Users/kmayer/Downloads/sorting-tests/1.digits"),
		tester.NewTester(&sortTask{fn: sort.SelectionSort[int]}, "/Users/kmayer/Downloads/sorting-tests/1.digits"),
		tester.NewTester(&sortTask{fn: sort.HeapSort[int]}, "/Users/kmayer/Downloads/sorting-tests/1.digits"),

		//tester.NewTester(&sortTask{fn: sort.BubbleSort[int]}, "/Users/kmayer/Downloads/sorting-tests/2.sorted"),
		//tester.NewTester(&sortTask{fn: sort.Insertion[int]}, "/Users/kmayer/Downloads/sorting-tests/2.sorted"),
		//tester.NewTester(&sortTask{fn: sort.InsertionShift[int]}, "/Users/kmayer/Downloads/sorting-tests/2.sorted"),
		//tester.NewTester(&sortTask{fn: sort.InsertionBinary[int]}, "/Users/kmayer/Downloads/sorting-tests/2.sorted"),
		//tester.NewTester(&sortTask{fn: sort.Shell[int]}, "/Users/kmayer/Downloads/sorting-tests/2.sorted"),
		tester.NewTester(&sortTask{fn: sort.SelectionSort[int]}, "/Users/kmayer/Downloads/sorting-tests/2.sorted"),
		tester.NewTester(&sortTask{fn: sort.HeapSort[int]}, "/Users/kmayer/Downloads/sorting-tests/2.sorted"),

		//tester.NewTester(&sortTask{fn: sort.BubbleSort[int]}, "/Users/kmayer/Downloads/sorting-tests/3.revers"),
		//tester.NewTester(&sortTask{fn: sort.Insertion[int]}, "/Users/kmayer/Downloads/sorting-tests/3.revers"),
		//tester.NewTester(&sortTask{fn: sort.InsertionShift[int]}, "/Users/kmayer/Downloads/sorting-tests/3.revers"),
		//tester.NewTester(&sortTask{fn: sort.InsertionBinary[int]}, "/Users/kmayer/Downloads/sorting-tests/3.revers"),
		//tester.NewTester(&sortTask{fn: sort.Shell[int]}, "/Users/kmayer/Downloads/sorting-tests/3.revers"),
		tester.NewTester(&sortTask{fn: sort.SelectionSort[int]}, "/Users/kmayer/Downloads/sorting-tests/3.revers"),
		tester.NewTester(&sortTask{fn: sort.HeapSort[int]}, "/Users/kmayer/Downloads/sorting-tests/3.revers"),
	}
	for _, t := range tests {
		t.RunTests()
	}
}

type sortTask struct {
	fn sort.Sorter[int]
}

func (t *sortTask) Run(args []string) (string, error) {
	if len(args) < 2 {
		panic("invalid arguments")
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("parse array length: %w", err)
	}

	arr := make([]int, 0, n)
	for _, s := range strings.Split(args[1], " ") {
		i, err := strconv.Atoi(s)
		if err != nil {
			return "", fmt.Errorf("parse number (%s): %w", s, err)
		}
		arr = append(arr, i)
	}

	_, err = withTimeout(t.fn, arr, getTimeoutForSize(len(arr)))
	if err != nil {
		return "", err
	}

	return sliceToString(arr), nil
}

func sliceToString(arr []int) string {
	builder := strings.Builder{}
	n := len(arr)
	for idx, i := range arr {
		builder.WriteString(strconv.Itoa(i))
		if idx < n-1 {
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}

func getTimeoutForSize(size int) time.Duration {
	timeout := 15 * time.Second
	if size > 100_000 {
		timeout = 2 * time.Minute
	}
	return timeout
}

func withTimeout(sorter sort.Sorter[int], arr []int, timeout time.Duration) (time.Duration, error) {
	ch := handle(sorter, arr)
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
