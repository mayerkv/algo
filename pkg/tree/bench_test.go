package tree

import (
	"math/rand"
	"testing"
)

type testCase struct {
	name       string
	createTree func() Tree
}

func BenchmarkTreeInsert(b *testing.B) {
	benchmarks := []testCase{
		{"Splay", NewSplayTree},
		{"Treap", NewTreap},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			tree := bm.createTree()
			rnd := newRand()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tree.Insert(rnd.Int())
			}
		})
	}
}

func BenchmarkTreeSearch(b *testing.B) {
	benchmarks := []testCase{
		{"Splay", NewSplayTree},
		{"Treap", NewTreap},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			tree := bm.createTree()
			rnd := newRand()
			insertRandom(tree, 1_000_000)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tree.Search(rnd.Intn(1_000_000))
			}
		})
	}
}

func BenchmarkTreeRemove(b *testing.B) {
	benchmarks := []testCase{
		{"Splay", NewSplayTree},
		{"Treap", NewTreap},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			tree := bm.createTree()
			rnd := newRand()
			insertRandom(tree, 1_000_000)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tree.Remove(rnd.Intn(1_000_000))
			}
		})
	}
}

func insertRandom(tree Tree, n int) {
	rnd := newRand()
	for i := 0; i < n; i++ {
		tree.Insert(rnd.Int())
	}
}

func newRand() *rand.Rand {
	return rand.New(rand.NewSource(0))
}
