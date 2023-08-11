package tester

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

type Tester struct {
	task Task
	path string
}

func NewTester(task Task, path string) *Tester {
	return &Tester{task: task, path: path}
}

func (t *Tester) RunTests() {
	nr := 0
	for {
		inFileName := fmt.Sprintf("%s/test.%d.in", t.path, nr)
		outFileName := fmt.Sprintf("%s/test.%d.out", t.path, nr)

		ok, err := allFilesExists(inFileName, outFileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !ok {
			return
		}

		result, err := t.RunTest(inFileName, outFileName)
		fmt.Printf("Test #%d: %v\n", nr, result)
		if err != nil {
			fmt.Println(err)
			return
		}

		nr++
	}
}

func allFilesExists(paths ...string) (bool, error) {
	for _, path := range paths {
		_, err := os.OpenFile(path, os.O_RDONLY, 0600)
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (t *Tester) RunTest(inFileName, outFileName string) (bool, error) {
	in, err := os.ReadFile(inFileName)
	if err != nil {
		return false, fmt.Errorf("cant read in file: %w", err)
	}
	out, err := os.ReadFile(outFileName)
	if err != nil {
		return false, fmt.Errorf("cant read out file: %w", err)
	}

	args := []string{string(in)}
	expected := strings.TrimSpace(string(out))
	actual, err := t.task.Run(args)
	if err != nil {
		return false, fmt.Errorf("run task error: %w", err)
	}

	return actual == expected, nil
}
