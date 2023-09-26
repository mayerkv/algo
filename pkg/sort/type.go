package sort

type Ordered interface {
	Int | String
}

type Int interface {
	~int
}

type String interface {
	~string
}

type Sorter[T Ordered] func(arr []T)
