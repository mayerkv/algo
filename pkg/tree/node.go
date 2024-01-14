package tree

type Node struct {
	key         int
	left, right *Node
}

func NewNode(key int) *Node {
	return &Node{key: key}
}

func rightRotate(x *Node) *Node {
	y := x.left
	x.left = y.right
	y.right = x
	return y
}

func leftRotate(x *Node) *Node {
	y := x.right
	x.right = y.left
	y.left = x
	return y
}

func (n *Node) ToOrderedSlice() []int {
	res := make([]int, 0)
	process := func(key int) {
		res = append(res, key)
	}
	preOrder(n, process)
	return res
}

func preOrder(n *Node, process func(int)) {
	if n == nil {
		return
	}
	preOrder(n.left, process)
	process(n.key)
	preOrder(n.right, process)
}
