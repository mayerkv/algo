package dynamic_array

type SingleArray[T any] struct {
	arr []T
}

func NewSingleArray[T any]() Array[T] {
	return &SingleArray[T]{
		arr: make([]T, 0, 0),
	}
}

func (a *SingleArray[T]) Add(item T, index int) {
	newSize := len(a.arr) + 1
	if index >= newSize {
		newSize = index + 1
	}
	arr := make([]T, newSize, newSize)
	pos := index
	if index > len(a.arr) {
		pos = len(a.arr)
	}
	copy(arr[0:pos], a.arr[0:pos])
	copy(arr[pos+1:], a.arr[pos:])
	arr[index] = item
	a.arr = arr
}

func (a *SingleArray[T]) Remove(index int) T {
	var noob T
	if index >= len(a.arr) {
		return noob
	}
	value := a.arr[index]
	lastIdx := len(a.arr) - 1
	copy(a.arr[index:], a.arr[index+1:])
	a.arr[lastIdx] = noob
	return value
}
