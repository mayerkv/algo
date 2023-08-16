package primer

import "math"

type Primer func(n int) int
type IsPrime func(p int) bool

// MakeSimplePrimer O(N) * O(IsPrime) = O(N^2)
func MakeSimplePrimer(isPrime IsPrime) Primer {
	return func(n int) int {
		var count int
		for p := 2; p <= n; p++ {
			if isPrime(p) {
				count++
			}
		}
		return count
	}
}

// IsPrimeA1 O(N)
func IsPrimeA1(p int) bool {
	count := 0
	for j := 1; j <= p; j++ {
		if p%j == 0 {
			count++
		}
	}

	return count == 2
}

// IsPrimeA2 O(N)
func IsPrimeA2(p int) bool {
	for j := 2; j < p; j++ {
		if p%j == 0 {
			return false
		}
	}
	return true
}

// IsPrimeA3 O(N/2) = O(N)
func IsPrimeA3(p int) bool {
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	for j := 3; j < p; j += 2 {
		if p%j == 0 {
			return false
		}
	}
	return true
}

// IsPrimeA4 O(N/2/2) = O(N)
func IsPrimeA4(p int) bool {
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	for j := 3; j < p/2; j += 2 {
		if p%j == 0 {
			return false
		}
	}
	return true
}

// IsPrimeA5 O(SQRT(N)/2) = O(SQRT(N))
func IsPrimeA5(p int) bool {
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	for j := 3; j <= int(math.Sqrt(float64(p))); j += 2 {
		if p%j == 0 {
			return false
		}
	}
	return true
}

// IsPrimeA6 O(SQRT(N)/2) = O(SQRT(N))
func IsPrimeA6(p int) bool {
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	x := int(math.Sqrt(float64(p)))
	for j := 3; j <= x; j += 2 {
		if p%j == 0 {
			return false
		}
	}
	return true
}

// Eratosthenes O(N Log Log N)
func Eratosthenes(n int) int {
	divs := make([]bool, n+1)
	count := 0
	sqrt := int(math.Sqrt(float64(n)))
	for p := 2; p <= n; p++ {
		if !divs[p] {
			count++
			if p <= sqrt {
				for j := p * p; j <= n; j += p {
					divs[j] = true
				}
			}
		}
	}
	return count
}
