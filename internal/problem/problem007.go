package problem

func problem007(params Params) (interface{}, error) {
	err := params.assertKeysExist("n")
	if err != nil {
		return nil, err
	}

	n, err := params.getInt("n")
	if err != nil {
		return nil, err
	}

	candidate := 1
	for i := 0; i < n; i++ {
		candidate++

		// calculate the next prime
		for ; !isPrime(candidate); candidate++ {
		}
	}

	return candidate, nil
}
