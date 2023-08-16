package main

import (
	"fmt"

	"github.com/mayerkv/algo/pkg/power"
)

func main() {
	a := 1.000_001
	n := 1_000_000

	for _, f := range []power.Power{power.Iterative, power.Recursive, power.Binary} {
		fmt.Println(f(a, n))
	}
}
