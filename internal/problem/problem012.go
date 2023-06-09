package problem

func problem012(params Params) (interface{}, error) {
	err := params.assertKeysExist("min")
	if err != nil {
		return nil, err
	}

	min, err := params.getInt("min")
	if err != nil {
		return nil, err
	}

	i := 1
	triangleNum := i
	factors := factorize(triangleNum)

	for len(factors) < min {
		i++
		triangleNum += i
		factors = factorize(triangleNum)
	}

	return triangleNum, nil
}

func factorize(n int) []int {
	if n == 0 {
		return []int{}
	}

	if n == 1 {
		return []int{1}
	}

	min := 1
	max := n
	factors := []int{}

	for min < max && min < (n/2)+1 {
		if n%min == 0 {
			max = n / min
			factors = append(factors, min, max)
		}
		min++
	}

	return factors
}
