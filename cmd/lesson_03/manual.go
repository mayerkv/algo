package main

import (
	"fmt"
	"time"

	"github.com/mayerkv/algo/pkg/fibonacci"
	"github.com/mayerkv/algo/pkg/power"
	"github.com/mayerkv/algo/pkg/primer"
)

func main() {
	runPower()
	runFibonacci()
	runPrimers()
}

func runPower() {
	fmt.Println("Run Power")

	a := 1.000_001
	n := 1_000_000
	pows := []power.Power{
		power.Iterative,
		power.Recursive,
		power.Binary,
	}
	for _, f := range pows {
		fmt.Println(f(a, n))
	}

	fmt.Println("Finish")
}

func runFibonacci() {
	fmt.Println("Run fibonacci")

	n := 20
	fibers := []fibonacci.Fibonacci{
		fibonacci.Iterative,
		fibonacci.Matrix,
		fibonacci.Recursive,
	}
	for _, f := range fibers {
		fmt.Println(f(n))
	}

	fmt.Println("Finish")
}

func runPrimers() {
	fmt.Println("Run primes")

	primers := []primer.Primer{
		primer.MakeSimplePrimer(primer.IsPrimeA1),
		primer.MakeSimplePrimer(primer.IsPrimeA2),
		primer.MakeSimplePrimer(primer.IsPrimeA3),
		primer.MakeSimplePrimer(primer.IsPrimeA4),
		primer.MakeSimplePrimer(primer.IsPrimeA5),
		primer.MakeSimplePrimer(primer.IsPrimeA6),
		primer.Eratosthenes,
	}

	for _, f := range primers {
		runPrimer(f)
	}

	fmt.Println("Finish")
}

func runPrimer(f primer.Primer) {
	defer fmt.Println("===")

	for n := 10; n <= 1_000_000_000; n *= 10 {
		ch := make(chan int, 1)
		timer := time.NewTimer(7 * time.Second)
		now := time.Now()
		go func() {
			defer close(ch)
			q := f(n)
			ch <- q
		}()
		select {
		case <-timer.C:
			fmt.Println("timeout")
			return
		case cnt := <-ch:
			fmt.Printf("%d -> %d %s\n", n, cnt, time.Since(now))
		}
	}
}
