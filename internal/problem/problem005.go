package problem

func problem005(params Params) (interface{}, error) {
	err := params.assertKeysExist("n")
	if err != nil {
		return nil, err
	}

	n, err := params.getInt("n")
	if err != nil {
		return nil, err
	}

	for i := n; true; i += n {
		if divisibleByAll(n, i) {
			return i, nil
		}
	}

	return nil, nil
}

func divisibleByAll(max int, x int) bool {
	for i := max; i > 0; i-- {
		if x%i != 0 {
			return false
		}
	}
	return true
}
