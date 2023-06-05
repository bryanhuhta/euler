package problem

import "math"

func problem003(params Params) (interface{}, error) {
	err := params.assertKeysExist("p")
	if err != nil {
		return nil, err
	}

	p, err := params.getInt("p")
	if err != nil {
		return nil, err
	}

	largestFactor := 0
	upperBound := int(math.Sqrt(float64(p))) + 1
	primes := firstNPrimes(upperBound)

	for _, prime := range primes {
		if p%prime == 0 {
			largestFactor = prime
		}
	}

	return largestFactor, nil
}

func firstNPrimes(n int) []int {
	primes := make([]int, 0, n)

	i := 2
	for len(primes) < n {
		if isPrime(i) {
			primes = append(primes, i)
		}
		i++
	}

	return primes
}
