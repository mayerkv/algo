package search

import "fmt"

type Node struct {
	key         int
	left, right *Node
	height      int
}

type Tree struct {
	root *Node
}

func (t *Tree) PreOrder() {
	preOrder(t.root)
	fmt.Println()
}

func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.key, " ")
	preOrder(n.left)
	preOrder(n.right)
}
