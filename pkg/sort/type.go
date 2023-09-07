package sort

type Ordered interface {
	~int | ~string
}

type Sorter[T Ordered] func(arr []T)
