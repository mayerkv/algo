package main

import (
	"fmt"
	"strconv"

	"github.com/mayerkv/algo/pkg/bits"
	"github.com/mayerkv/algo/pkg/tester"
)

func main() {
	tests := []*tester.Tester{
		tester.NewTester(&chessTask{bits.King, bits.PopCnt3}, "/Users/kmayer/Downloads/0.BITS/1.Bitboard - Король"),
		tester.NewTester(&chessTask{bits.Knight, bits.PopCnt3}, "/Users/kmayer/Downloads/0.BITS/2.Bitboard - Конь"),
		tester.NewTester(&chessTask{bits.Rook, bits.PopCnt3}, "/Users/kmayer/Downloads/0.BITS/3.Bitboard - Ладья"),
	}

	for _, t := range tests {
		t.RunTests()
	}
}

type chessTask struct {
	bitsFn func(n int) uint64
	cntFn  func(b uint64) int
}

func (t *chessTask) Run(args []string) (string, error) {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return "", err
	}
	b := t.bitsFn(n)
	cnt := t.cntFn(b)

	return fmt.Sprintf("%d\r\n%d", cnt, b), nil
}
