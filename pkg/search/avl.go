package search

type AVL struct {
	BST
}

func NewAVL() *AVL {
	return &AVL{}
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(x int) *Node {
	return &Node{
		key:    x,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func rightRotate(n *Node) *Node {
	x := n.left
	y := x.right

	x.right = n
	n.left = y

	n.height = max(height(n.left), height(n.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1
	return x
}

func leftRotate(n *Node) *Node {
	x := n.right
	y := x.left

	x.left = n
	n.right = y

	n.height = max(height(n.left), height(n.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func (t *AVL) Insert(x int) {
	t.root = avlInsert(t.root, x)
}

func avlInsert(n *Node, key int) *Node {
	if n == nil {
		return newNode(key)
	}
	if key < n.key {
		n.left = avlInsert(n.left, key)
	} else if key > n.key {
		n.right = avlInsert(n.right, key)
	} else {
		return n
	}
	n.height = max(height(n.left), height(n.right)) + 1

	balance := getBalance(n)
	if balance > 1 && key < n.left.key {
		return rightRotate(n)
	}
	if balance < -1 && key > n.right.key {
		return leftRotate(n)
	}
	if balance > 1 && key > n.left.key {
		n.left = leftRotate(n.left)
		return rightRotate(n)
	}
	if balance < -1 && key < n.right.key {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}
	return n
}

func minValueNode(n *Node) *Node {
	curr := n
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (t *AVL) Delete(x int) {
	t.root = avlDeleteNode(t.root, x)
}

func avlDeleteNode(n *Node, key int) *Node {
	if n == nil {
		return n
	}
	if key < n.key {
		n.left = avlDeleteNode(n.left, key)
	} else if key > n.key {
		n.right = avlDeleteNode(n.right, key)
	} else {
		if n.left == nil || n.right == nil {
			var tmp *Node
			if n.left != nil {
				tmp = n.left
			} else {
				tmp = n.right
			}
			if tmp == nil {
				tmp = n
				n = nil
			} else {
				*n = *tmp
			}
		} else {
			tmp := minValueNode(n.right)
			n.key = tmp.key
			n.right = avlDeleteNode(n.right, tmp.key)
		}
	}
	if n == nil {
		return n
	}

	n.height = max(height(n.left), height(n.right)) + 1
	balance := getBalance(n)

	if balance > 1 && getBalance(n.left) >= 0 {
		return rightRotate(n)
	}
	if balance > 1 && getBalance(n.left) < 0 {
		n.left = leftRotate(n.left)
		return rightRotate(n)
	}
	if balance < -1 && getBalance(n.right) <= 0 {
		return leftRotate(n)
	}
	if balance < -1 && getBalance(n.right) > 0 {
		n.right = rightRotate(n.right)
		return leftRotate(n)
	}
	return n
}
