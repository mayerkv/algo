package main

import (
	"fmt"
	"github.com/mayerkv/algo/pkg/tree"
	"math/rand"
	"time"
)

func main() {
	n := 1000

	fmt.Println("Splay Tree")
	test("SEARCH RANDOM 100", testSearchRandom100, createRandomN(n), createOrderedN(n))
	test("SEARCH 1-10  1000", testSearch1000from1to10, createRandomN(n), createOrderedN(n))
	test("REMOVE RANDOM 100", testRemoveRandom100, createRandomN(n), createOrderedN(n))
}

func test(name string, fn testFn, rnd, ord *tree.SplayTree) {
	fmt.Printf("[TEST %s] random = %d, ordered = %d\n", name, fn(rnd), fn(ord))
}

type testFn func(tree *tree.SplayTree) time.Duration

func testSearchRandom100(tree *tree.SplayTree) time.Duration {
	rnd := createRnd()
	now := time.Now()
	for i := 0; i < 100; i++ {
		tree.Search(rnd.Intn(1000))
	}
	return time.Now().Sub(now)
}

func testSearch1000from1to10(tree *tree.SplayTree) time.Duration {
	now := time.Now()
	for i := 0; i < 1000; i++ {
		for j := 1; j <= 10; j++ {
			tree.Search(j)
		}
	}
	return time.Now().Sub(now)
}

func testRemoveRandom100(tree *tree.SplayTree) time.Duration {
	rnd := createRnd()
	now := time.Now()
	for i := 0; i < 100; i++ {
		tree.Remove(rnd.Intn(1000))
	}
	return time.Now().Sub(now)
}

func createRnd() *rand.Rand {
	return rand.New(rand.NewSource(0))
}

func createRandomN(n int) *tree.SplayTree {
	t := tree.NewSplayTree()
	rnd := createRnd()
	for i := 0; i < n; i++ {
		t.Insert(rnd.Intn(n))
	}
	return t
}

// Вставка в порядке возрастания осуществляется быстрее.
// Новый элемент идет направо от рута, выполняется левый поворот
func createOrderedN(n int) *tree.SplayTree {
	t := tree.NewSplayTree()
	for i := 0; i < n; i++ {
		t.Insert(i)
	}
	return t
}
