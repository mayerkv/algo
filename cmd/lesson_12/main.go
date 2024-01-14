package main

import (
	"fmt"
	"github.com/mayerkv/algo/pkg/tree"
	"math/rand"
	"time"
)

func main() {
	n := 1000
	var tests = []struct {
		name    string
		factory func() tree.Tree
	}{
		{"Splay Tree", tree.NewSplayTree},
		{"Treap", tree.NewTreap},
	}
	var cases = []struct {
		name string
		test func(tree tree.Tree) time.Duration
	}{
		{"SEARCH RANDOM 100", testSearchRandom100},
		{"SEARCH 1-10  1000", testSearch1000from1to10},
		{"REMOVE RANDOM 100", testRemoveRandom100},
	}

	for _, t := range tests {
		fmt.Println(t.name)
		for _, c := range cases {
			test(c.name, c.test, createRandomN(t.factory, n), createOrderedN(t.factory, n))
		}
	}
}

func test(name string, fn testFn, rnd, ord tree.Tree) {
	fmt.Printf("[%s] random = %s ordered = %s\n", name, fn(rnd), fn(ord))
}

type testFn func(tree tree.Tree) time.Duration

func testSearchRandom100(tree tree.Tree) time.Duration {
	rnd := createRnd()
	now := time.Now()
	for i := 0; i < 100; i++ {
		tree.Search(rnd.Intn(1000))
	}
	return time.Now().Sub(now)
}

func testSearch1000from1to10(tree tree.Tree) time.Duration {
	now := time.Now()
	for i := 0; i < 1000; i++ {
		for j := 1; j <= 10; j++ {
			tree.Search(j)
		}
	}
	return time.Now().Sub(now)
}

func testRemoveRandom100(tree tree.Tree) time.Duration {
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

func createRandomN(factory func() tree.Tree, n int) tree.Tree {
	rnd := createRnd()
	t := factory()
	for i := 0; i < n; i++ {
		t.Insert(rnd.Intn(n))
	}
	return t
}

func createOrderedN(factory func() tree.Tree, n int) tree.Tree {
	t := factory()
	for i := 0; i < n; i++ {
		t.Insert(i)
	}
	return t
}
