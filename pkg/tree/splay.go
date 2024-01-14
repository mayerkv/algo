package tree

type SplayTree struct {
	root *Node
}

func NewSplayTree() *SplayTree {
	return &SplayTree{}
}

func (t *SplayTree) Insert(key int) {
	t.root = splayInsert(t.root, key)
}

func (t *SplayTree) Search(x int) bool {
	node := splay(t.root, x)
	if node == nil {
		return false
	}
	t.root = node
	return true
}

func (t *SplayTree) Remove(x int) {
	t.root = splayRemove(t.root, x)
}

func splay(root *Node, key int) *Node {
	if root == nil || root.key == key {
		return root
	}
	if root.key > key {
		if root.left == nil {
			return root
		}
		if root.left.key > key {
			root.left.left = splay(root.left.left, key)
			root = rightRotate(root)
		} else if root.left.key < key {
			root.left.right = splay(root.left.right, key)
			if root.left.right != nil {
				root.left = leftRotate(root.left)
			}
		}
		if root.left == nil {
			return root
		}
		return rightRotate(root)
	}
	if root.right == nil {
		return root
	}
	if root.right.key > key {
		root.right.left = splay(root.right.left, key)
		if root.right.left != nil {
			root.right = rightRotate(root.right)
		}
	} else if root.right.key < key {
		root.right.right = splay(root.right.right, key)
		root = leftRotate(root)
	}
	if root.right == nil {
		return root
	}
	return leftRotate(root)
}

func splayInsert(root *Node, key int) *Node {
	if root == nil {
		return NewNode(key)
	}
	root = splay(root, key)
	if root.key == key {
		return root
	}
	tmp := NewNode(key)
	if root.key > key {
		tmp.right = root
		tmp.left = root.left
		root.left = nil
	} else {
		tmp.left = root
		tmp.right = root.right
		root.right = nil
	}
	return tmp
}

func splayRemove(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	root = splay(root, key)
	if key != root.key {
		return root
	}
	tmp := root
	if root.left == nil {
		root = root.right
	} else {
		root = splay(root.left, key)
		root.right = tmp.right
	}
	return root
}
