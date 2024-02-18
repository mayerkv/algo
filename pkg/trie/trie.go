package trie

const alphabetSize = 128

type Trie struct {
	root *trieNode
}

func Constructor() Trie {
	return Trie{
		root: newNode(),
	}
}

func (t *Trie) Insert(word string) {
	n := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if n.children[idx] == nil {
			n.children[idx] = newNode()
		}
		n = n.children[idx]
	}
	n.isFinal = true
}

func (t *Trie) Search(word string) bool {
	n := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if n.children[idx] == nil {
			return false
		}
		n = n.children[idx]
	}
	return n.isFinal
}

func (t *Trie) StartsWith(prefix string) bool {
	n := t.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if n.children[idx] == nil {
			return false
		}
		n = n.children[idx]
	}
	return true
}

type trieNode struct {
	children [alphabetSize]*trieNode
	isFinal  bool
}

func newNode() *trieNode {
	return &trieNode{}
}
