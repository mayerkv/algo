package tester

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

const (
	MaxBufferSize         = 1024 * 1024 * 80 // 80 Mb
	InitialBufferCapacity = 64 * 1024        // 64 Kb
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
		if err != nil {
			fmt.Printf("Test #%d error: %s\n", nr, err.Error())
			return
		}
		fmt.Printf("Test #%d: %v\n", nr, result)

		nr++
	}
}

func allFilesExists(paths ...string) (bool, error) {
	for _, path := range paths {
		f, err := os.OpenFile(path, os.O_RDONLY, 0600)
		f.Close()
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
	in, err := os.OpenFile(inFileName, os.O_RDONLY, 0600)
	if err != nil {
		return false, fmt.Errorf("cant read in file: %w", err)
	}
	defer in.Close()

	out, err := os.ReadFile(outFileName)
	if err != nil {
		return false, fmt.Errorf("cant read out file: %w", err)
	}

	var args []string
	fileScanner := bufio.NewScanner(in)
	fileScanner.Split(bufio.ScanLines)
	buf := make([]byte, 0, InitialBufferCapacity)
	fileScanner.Buffer(buf, MaxBufferSize)

	for fileScanner.Scan() {
		args = append(args, fileScanner.Text())
	}

	expected := strings.TrimSpace(string(out))
	actual, err := t.task.Run(args)
	if err != nil {
		return false, fmt.Errorf("run task error: %w", err)
	}

	return actual == expected, nil
}
