package problem

func problem002(params Params) (interface{}, error) {
	err := params.assertKeysExist("max")
	if err != nil {
		return nil, err
	}

	max, err := params.getInt("max")
	if err != nil {
		return nil, err
	}

	a := 1
	b := 2
	c := a + b

	sum := 2 // Count the initial 2 value
	for c < max {
		if c%2 == 0 {
			sum += c
		}

		a = b
		b = c
		c = a + b
	}

	return sum, nil
}
