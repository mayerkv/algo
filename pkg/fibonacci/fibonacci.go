package fibonacci

type Fibonacci func(n int) int

// Recursive O(2^N)
func Recursive(n int) int {
	if n <= 1 {
		return n
	}
	return Recursive(n-1) + Recursive(n-2)
}

// Iterative O(N)
func Iterative(n int) int {
	last, next := 0, 1
	for i := 0; i < n; i++ {
		last, next = next, last+next
	}
	return last
}

// Matrix O(log N)
func Matrix(n int) int {
	m := BASE.pow(n - 1)
	return m.a
}
