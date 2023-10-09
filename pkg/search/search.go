package search

type Search interface {
	Insert(x int)
	Search(x int) bool
	Remove(x int)
}
