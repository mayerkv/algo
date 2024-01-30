package tree

import "math/rand"

type treapNode struct {
	key         int
	priority    int
	left, right *treapNode
}

func newTreapNode(key int) *treapNode {
	return &treapNode{key: key, priority: rand.Int()}
}

type Treap struct {
	root *treapNode
}

func NewTreap() Tree {
	return &Treap{}
}

func (t *Treap) Insert(x int) {
	t1, t2 := treapSplit(t.root, x)
	n := newTreapNode(x)
	t.root = treapMerge(treapMerge(t1, n), t2)
}

func (t *Treap) Search(x int) bool {
	var curr = t.root
	for curr != nil {
		if curr.key == x {
			return true
		} else if curr.key > x {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (t *Treap) Remove(x int) {
	t.root = treapRemove(t.root, x)
}

func treapInsert(root *treapNode, key int) *treapNode {
	if root == nil {
		return newTreapNode(key)
	}
	if key < root.key {
		root.left = treapInsert(root.left, key)
		if root.left.priority > root.priority {
			root = treapRightRotate(root)
		}
	} else {
		root.right = treapInsert(root.right, key)
		if root.right.priority > root.priority {
			root = treapLeftRotate(root)
		}
	}
	return root
}

func treapRemove(root *treapNode, key int) *treapNode {
	if root == nil {
		return root
	}
	if key < root.key {
		root.left = treapRemove(root.left, key)
	}
	if key > root.key {
		root.right = treapRemove(root.right, key)
	} else if root.left == nil {
		root = root.right
	} else if root.right == nil {
		root = root.left
	} else if root.left.priority < root.right.priority {
		root = treapLeftRotate(root)
		root.left = treapRemove(root.left, key)
	} else {
		root = treapRightRotate(root)
		root.right = treapRemove(root.right, key)
	}
	return root
}

func treapRightRotate(x *treapNode) *treapNode {
	y := x.left
	x.left = y.right
	y.right = x
	return y
}

func treapLeftRotate(x *treapNode) *treapNode {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

func treapMerge(t1, t2 *treapNode) *treapNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	if t1.priority > t2.priority {
		t1.right = treapMerge(t1.right, t2)
		return t1
	}
	t2.left = treapMerge(t1, t2.left)
	return t2
}

func treapSplit(n *treapNode, key int) (*treapNode, *treapNode) {
	if n == nil {
		return nil, nil
	}
	if key > n.key {
		t1, t2 := treapSplit(n.right, key)
		n.right = t1
		return n, t2
	}
	t1, t2 := treapSplit(n.left, key)
	n.left = t2
	return t1, n
}
