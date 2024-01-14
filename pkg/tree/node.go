package tree

type node struct {
	key         int
	left, right *node
}

func newNode(key int) *node {
	return &node{key: key}
}

func rightRotate(x *node) *node {
	y := x.left
	x.left = y.right
	y.right = x
	return y
}

func leftRotate(x *node) *node {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

func (n *node) ToOrderedSlice() []int {
	res := make([]int, 0)
	process := func(key int) {
		res = append(res, key)
	}
	preOrder(n, process)
	return res
}

func preOrder(n *node, process func(int)) {
	if n == nil {
		return
	}
	preOrder(n.left, process)
	process(n.key)
	preOrder(n.right, process)
}
