package dynamic_array

type FactorArray[T any] struct {
	arr    []T
	factor int
	size   int
}

func NewFactorArray[T any](factor, initLength int) Array[T] {
	return &FactorArray[T]{
		arr:    make([]T, initLength, initLength),
		factor: factor,
		size:   0,
	}
}

func (a *FactorArray[T]) Add(item T, index int) {
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

func (a *FactorArray[T]) Remove(index int) T {
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

func (a *FactorArray[T]) allocForIndex(index int) []T {
	capacity := len(a.arr)
	newSize := a.size + 1
	if index > capacity {
		newSize = index + 1
	}
	if newSize <= capacity {
		return a.arr
	}
	for newSize > capacity {
		capacity += capacity * a.factor / 100
	}
	return make([]T, capacity, capacity)
}
