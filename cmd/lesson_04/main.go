package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/mayerkv/algo/pkg/dynamic_array"
)

func main() {
	for i := 100; i <= 100_000; i *= 10 {
		arrays := []dynamic_array.Array[int]{
			dynamic_array.NewSingleArray[int](),
			dynamic_array.NewVectorArray[int](10),
			dynamic_array.NewFactorArray[int](50, 100),
		}

		for _, array := range arrays {
			testAdd(array, i)
			testRemove(array, i)
			fmt.Println()
		}

	}
}

func testAdd(array dynamic_array.Array[int], n int) {
	now := time.Now()
	for i := 0; i < n; i++ {
		array.Add(n+1, n-i-1)
	}
	log.Printf("%v.testAdd, n = %d %s", reflect.TypeOf(array), n, time.Since(now))
}

func testRemove(array dynamic_array.Array[int], n int) {
	now := time.Now()
	for i := 0; i < n; i++ {
		array.Remove(n - i)
	}
	log.Printf("%v.testRemove, n = %d %s", reflect.TypeOf(array), n, time.Since(now))
}
