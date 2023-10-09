package search

type BST struct {
	Tree
}

func NewBST() *BST {
	return &BST{
		Tree{root: nil},
	}
}

func CreateBST(arr []int) *BST {
	return &BST{
		Tree: Tree{
			root: createBST(arr, 0, len(arr)-1),
		},
	}
}

func (t *BST) Insert(x int) {
	t.root = btsInsert(t.root, x)
}

func (t *BST) Search(x int) bool {
	var curr = t.root
	for curr != nil {
		if curr.value == x {
			return true
		} else if curr.value > x {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}

func (t *BST) Remove(x int) {
	t.root = bstRemove(t.root, x)
}

func createBST(arr []int, start, end int) *Node {
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	return &Node{
		value: arr[mid],
		left:  createBST(arr, start, mid-1),
		right: createBST(arr, mid+1, end),
	}
}

func btsInsert(n *Node, x int) *Node {
	if n == nil {
		return &Node{value: x}
	}
	if x < n.value {
		n.left = btsInsert(n.left, x)
	} else {
		n.right = btsInsert(n.right, x)
	}
	return n
}

func bstRemove(n *Node, x int) *Node {
	if n != nil {
		if n.value == x {
			if n.left == nil && n.right == nil {
				return nil
			}
			if n.left == nil {
				return n.right
			}
			if n.right == nil {
				return n.left
			}
			maxNode := bstFindMax(n.left)
			n.value = maxNode.value
			n.left = bstRemove(n.left, maxNode.value)
		} else {
			if n.value > x {
				n.left = bstRemove(n.left, x)
			} else {
				n.right = bstRemove(n.right, x)
			}
		}
	}
	return n
}

func bstFindMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}
