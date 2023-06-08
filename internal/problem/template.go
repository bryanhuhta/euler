package problem

func problemXXX(params Params) (interface{}, error) {
	err := params.assertKeysExist("X")
	if err != nil {
		return nil, err
	}

	x, err := params.getInt("x")
	if err != nil {
		return nil, err
	}

	return x, nil
}
