package main

import (
	"os"

	"github.com/mayerkv/algo/pkg/sort"
)

func main() {
	input := "in.txt"
	output := "out.txt"
	generateInput(input, 50, 100)
	err := mergeSort(input, output)
	checkError(err)
}

func mergeSort(input, output string) error {
	in := openFile(input, os.O_RDONLY)
	defer in.Close()

	out := openFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR)
	defer out.Close()

	chunks, err := splitChunks[int](in, 100, sort.MergeSort[int])
	if err != nil {
		return err
	}

	return mergeChunks[int](chunks, out)
}
