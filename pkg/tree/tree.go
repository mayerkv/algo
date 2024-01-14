package tree

type Tree interface {
	Insert(key int)
	Search(key int) bool
	Remove(key int)
}
