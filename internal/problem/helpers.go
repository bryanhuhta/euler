package problem

import "math"

func isPrime(n int) bool {
	// Sqrt, then floor, then add 1 to simulate rounding up.
	max := int(math.Sqrt(float64(n))) + 1

	for i := 2; i < max; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
