package problem

func problem006(params Params) (interface{}, error) {
	err := params.assertKeysExist("n")
	if err != nil {
		return nil, err
	}

	n, err := params.getInt("n")
	if err != nil {
		return nil, err
	}

	return squareOfSum(n) - sumOfSquares(n), nil
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}
