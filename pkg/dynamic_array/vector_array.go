package dynamic_array

type VectorArray[T any] struct {
	arr    []T
	size   int
	vector int
}

func NewVectorArray[T any](vector int) Array[T] {
	return &VectorArray[T]{
		arr:    make([]T, 0, 0),
		size:   0,
		vector: vector,
	}
}

func (a *VectorArray[T]) Add(item T, index int) {
	arr := a.allocForIndex(index)
	pos := index
	if index > len(a.arr) {
		pos = len(a.arr)
	}
	copy(arr[0:pos], a.arr[0:pos])
	copy(arr[pos+1:], a.arr[pos:])
	arr[index] = item
	a.arr = arr
	a.size++
}

func (a *VectorArray[T]) Remove(index int) T {
	var noob T
	lastIdx := a.size - 1
	if index > lastIdx {
		return noob
	}
	value := a.arr[index]
	copy(a.arr[index:], a.arr[index+1:])
	a.arr[lastIdx] = noob
	a.size--
	return value
}

func (a *VectorArray[T]) allocForIndex(index int) []T {
	capacity := len(a.arr)
	newSize := a.size + 1
	if index > capacity {
		newSize = index + 1
	}
	if newSize <= capacity {
		return a.arr
	}
	for newSize > capacity {
		capacity += a.vector
	}
	return make([]T, capacity, capacity)
}
