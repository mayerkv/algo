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
	//TODO implement me
	panic("implement me")
}

func (a *FactorArray[T]) Remove(index int) T {
	//TODO implement me
	panic("implement me")
}
