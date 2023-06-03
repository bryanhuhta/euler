package problem

func problem001(params Params) (interface{}, error) {
	err := params.assertKeysExist("max")
	if err != nil {
		return nil, err
	}

	max, err := params.getInt("max")
	if err != nil {
		return nil, err
	}

	sum := 0
	for i := 0; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum, nil
}
