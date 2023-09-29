package sort

func BucketSort[T Int](arr []T) {
	n := len(arr)
	m := max(arr) + 1

	bucket := make([]*list[T], n)
	for _, value := range arr {
		nr := value * T(n) / m
		if bucket[nr] == nil {
			bucket[nr] = &list[T]{}
		}
		bucket[nr].addSort(value)
	}

	var j int
	for _, b := range bucket {
		if b == nil {
			continue
		}
		iter := b.iterator()
		for iter.HasNext() {
			arr[j] = iter.Value()
			j++
		}
	}
}

func max[T Int](arr []T) T {
	var v T
	for _, i := range arr {
		if i > v {
			v = i
		}
	}
	return v
}

type node[T Int] struct {
	value T
	next  *node[T]
}

type list[T Int] struct {
	head *node[T]
}

func (l *list[T]) addSort(value T) {
	var prev *node[T]
	curr := l.head
	n := &node[T]{value, nil}
	for curr != nil && n.value > curr.value {
		prev, curr = curr, curr.next
	}
	if prev == nil {
		l.head = n
	} else {
		prev.next = n
	}
	n.next = curr
}

func (l *list[T]) iterator() *iterator[T] {
	return &iterator[T]{pos: l.head}
}

type iterator[T Int] struct {
	pos *node[T]
}

func (i *iterator[T]) HasNext() bool {
	return i.pos != nil
}

func (i *iterator[T]) Value() T {
	if !i.HasNext() {
		var noop T
		return noop
	}
	value := i.pos.value
	i.pos = i.pos.next
	return value
}
