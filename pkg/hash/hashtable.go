package hash

const (
	defaultCapacity = 11
	loadFactor      = .75
)

type Key interface {
	comparable
}

type HashFunc[T Key] func(T) int

type hashEntry[K Key, V any] struct {
	key   K
	value V
	next  *hashEntry[K, V]
}

type HashTable[K Key, V any] struct {
	size      int
	threshold int
	hashFunc  HashFunc[K]
	buckets   []*hashEntry[K, V]
}

func NewHashTable[K Key, V any](hashFunc HashFunc[K]) *HashTable[K, V] {
	t := float64(defaultCapacity) * loadFactor
	return &HashTable[K, V]{
		size:      0,
		hashFunc:  hashFunc,
		buckets:   make([]*hashEntry[K, V], defaultCapacity),
		threshold: int(t),
	}
}

func (t *HashTable[K, V]) Size() int {
	return t.size
}

func (t *HashTable[K, V]) IsEmpty() bool {
	return t.size == 0
}

func (t *HashTable[K, V]) Get(key K) (V, bool) {
	idx := t.hash(key)
	e := t.buckets[idx]
	for e != nil {
		if e.key == key {
			return e.value, true
		}
		e = e.next
	}
	var noop V
	return noop, false
}

func (t *HashTable[K, V]) Put(key K, value V) {
	idx := t.hash(key)
	e := t.buckets[idx]
	for e != nil {
		if e.key == key {
			e.value = value
			return
		}
		e = e.next
	}
	t.size++
	if t.size > t.threshold {
		t.rehash()
		idx = t.hash(key)
	}
	t.buckets[idx] = &hashEntry[K, V]{key, value, t.buckets[idx]}
}

func (t *HashTable[K, V]) Remove(key K) {
	idx := t.hash(key)
	e := t.buckets[idx]
	var last *hashEntry[K, V]
	for e != nil {
		if e.key == key {
			if last == nil {
				t.buckets[idx] = e.next
			} else {
				last.next = e.next
			}
			t.size--
			return
		}
		last = e
		e = e.next
	}
}

func (t *HashTable[K, V]) hash(key K) int {
	sum := t.hashFunc(key) % len(t.buckets)
	if sum < 0 {
		return -sum
	}
	return sum
}

func (t *HashTable[K, V]) rehash() {
	oldBuckets := t.buckets
	capacity := float64(len(t.buckets)*2 + 1)
	t.threshold = int(capacity * loadFactor)
	t.buckets = make([]*hashEntry[K, V], int(capacity))
	for i := len(oldBuckets) - 1; i >= 0; i-- {
		e := oldBuckets[i]
		for e != nil {
			idx := t.hash(e.key)
			dest := t.buckets[idx]
			if dest != nil {
				next := dest.next
				for next != nil {
					dest, next = next, dest.next
				}
				dest.next = e
			} else {
				t.buckets[idx] = e
			}
			next := e.next
			e.next = nil
			e = next
		}
	}
}
