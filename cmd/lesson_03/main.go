package main

import (
	"fmt"

	"github.com/mayerkv/algo/pkg/fibonacci"
	"github.com/mayerkv/algo/pkg/power"
)

func main() {
	//runPower()
	runFibonacci()
}

func runPower() {
	a := 1.000_001
	n := 1_000_000

	for _, f := range []power.Power{power.Iterative, power.Recursive, power.Binary} {
		fmt.Println(f(a, n))
	}
}

func runFibonacci() {
	n := 20

	for _, f := range []fibonacci.Fibonacci{fibonacci.Iterative, fibonacci.Matrix, fibonacci.Recursive} {
		fmt.Println(f(n))
	}
}
