package main

import (
	"flag"
	"strconv"

	"github.com/mayerkv/algo/pkg/lucky"
	"github.com/mayerkv/algo/pkg/tester"
)

// решение из статьи https://habr.com/ru/articles/266479/
func main() {
	var path string
	flag.StringVar(&path, "path", "/tmp/tests", "Path to testing directory")
	flag.Parse()

	task := newLuckyTask()
	tr := tester.NewTester(task, path)
	tr.RunTests()
}

type luckyTask struct {
}

func newLuckyTask() *luckyTask {
	return &luckyTask{}
}

func (t *luckyTask) Run(args []string) (string, error) {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return "", err
	}

	count := lucky.LuckyTicketsCount(n)

	return strconv.Itoa(count), nil
}
