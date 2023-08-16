package fibonacci

var (
	BASE = matrix{
		1, 1,
		1, 0,
	}
	IDENTITY = matrix{
		1, 0,
		0, 1,
	}
)

type matrix struct {
	a, b int
	c, d int
}

func (m1 matrix) multiply(m2 matrix) matrix {
	return matrix{
		a: m1.a*m2.a + m1.b*m2.c, b: m1.a*m2.b + m1.b*m2.d,
		c: m1.c*m2.a + m1.d*m2.c, d: m1.c*m2.b + m1.d*m2.d,
	}
}

func (m1 matrix) pow(n int) matrix {
	if n == 0 {
		return IDENTITY
	}
	if n == 1 {
		return m1
	}
	if n%2 == 0 {
		r := m1.pow(n / 2)
		return r.multiply(r)
	}

	return m1.pow(n - 1).multiply(m1)
}
