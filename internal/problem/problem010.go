package problem

func problem010(params Params) (interface{}, error) {
	err := params.assertKeysExist("max")
	if err != nil {
		return nil, err
	}

	max, err := params.getInt("max")
	if err != nil {
		return nil, err
	}

	sum := 0
	for i := 2; i < max; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	return sum, nil
}
