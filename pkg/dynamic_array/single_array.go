package dynamic_array

type SingleArray[T any] struct {
	arr []T
}

func NewSingleArray[T any]() *SingleArray[T] {
	return &SingleArray[T]{
		arr: make([]T, 0, 0),
	}
}

func (a *SingleArray[T]) Add(item T, index int) {
	capacity := cap(a.arr)
	if index > capacity-1 {
		newArr := make([]T, index+1)
		for idx, i := range a.arr {
			newArr[idx] = i
		}
		a.arr = newArr
	}
	a.arr[index] = item
}

func (a *SingleArray[T]) Remove(index int) T {
	capacity := cap(a.arr)
	if index > capacity-1 {
		panic("out of range")
	}
	item := a.arr[index]
	newArr := make([]T, capacity-1)
	for idx, i := range a.arr {
		if idx == index {
			continue
		}
		if idx < index {
			newArr[idx] = i
		} else {
			newArr[idx-1] = i
		}
	}
	a.arr = newArr

	return item
}
