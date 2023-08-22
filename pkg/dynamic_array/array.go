package dynamic_array

type Array[T any] interface {
	Add(item T, index int)
	Remove(index int) T
}
