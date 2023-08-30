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
	for idx := range a.arr {
		if idx < index {
			arr[idx] = a.arr[idx]
		} else {
			arr[idx+1] = a.arr[idx]
		}
	}
	arr[index] = item
	a.arr = arr
}

func (a *SingleArray[T]) Remove(index int) (value T) {
	if index >= len(a.arr) {
		return value
	}
	value = a.arr[index]
	var noop T
	a.arr[index] = noop
	for idx := range a.arr {
		if idx > index {
			a.arr[idx-1], a.arr[idx] = a.arr[idx], a.arr[idx-1]
		}
	}
	return value
}
