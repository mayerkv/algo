package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mayerkv/algo/pkg/search"
)

func main() {
	//n := 10_000_000
	//testRandom(n)
	//testOrdered(n)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bst := search.CreateBST(arr)
	bst.PreOrder()
}

func testRandom(max int) {
	fmt.Println("Test random")
	for n := 1_000; n <= max; n *= 10 {
		bst := search.NewAVL()
		testInsert(newRandomIterator(n), bst)
		testSearch(newRandomIterator(n/10), bst)
		testRemove(newRandomIterator(n/10), bst)
	}
}

func testOrdered(max int) {
	fmt.Println("Test asc ordered")
	bst := search.NewAVL()
	for n := 1_000; n <= max; n *= 10 {
		testInsert(newAscOrderIterator(n), bst)
		testSearch(newAscOrderIterator(n/10), bst)
		testRemove(newAscOrderIterator(n/10), bst)
	}
}

func testInsert(it iterator, search search.Search) {
	now := time.Now()
	cnt := 0
	for it.Next() {
		search.Insert(it.Value())
		cnt++
	}
	fmt.Printf("insert %d items at %s\n", cnt, time.Since(now))
}

func testSearch(it iterator, search search.Search) {
	now := time.Now()
	cnt := 0
	for it.Next() {
		search.Search(it.Value())
		cnt++
	}
	fmt.Printf("search %d items at %s\n", cnt, time.Since(now))
}

func testRemove(it iterator, search search.Search) {
	now := time.Now()
	cnt := 0
	for it.Next() {
		search.Remove(it.Value())
		cnt++
	}
	fmt.Printf("remove %d items at %s\n", cnt, time.Since(now))
}

type iterator interface {
	Next() bool
	Value() int
}

type randomIterator struct {
	n   int
	rnd *rand.Rand
	v   int
}

func newRandomIterator(n int) iterator {
	return &randomIterator{
		n:   n,
		rnd: randomizer(123),
		v:   0,
	}
}

func (i *randomIterator) Next() bool {
	if i.n > 0 {
		i.v = i.rnd.Int()
		i.n--
		return true
	}
	return false
}

func (i *randomIterator) Value() int {
	return i.v
}

func randomizer(seed int64) *rand.Rand {
	src := rand.NewSource(seed)
	return rand.New(src)
}

type ascOrderIterator struct {
	n int
	v int
}

func newAscOrderIterator(n int) iterator {
	return &ascOrderIterator{n: n}
}

func (it *ascOrderIterator) Next() bool {
	if it.n > it.v {
		it.v++
		return true
	}
	return false
}

func (it *ascOrderIterator) Value() int {
	return it.v
}
