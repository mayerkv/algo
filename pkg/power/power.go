package power

type Power func(a float64, n int) float64

func Iterative(a float64, n int) float64 {
	res := 1.0
	for i := 0; i < n; i++ {
		res = res * a
	}
	return res
}

func Recursive(a float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	return Recursive(a, n-1) * a
}

func Binary(a float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return a
	}
	if n%2 == 0 {
		r := Binary(a, n/2)
		return r * r
	}
	return a * Binary(a, n-1)
}
