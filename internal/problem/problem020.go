package problem

func problem020(params Params) (interface{}, error) {
	err := params.assertKeysExist("n")
	if err != nil {
		return nil, err
	}

	n, err := params.getInt("n")
	if err != nil {
		return nil, err
	}

	factorial := []int{1}
	for i := 2; i <= n; i++ {
		factorial = bigMultiply(i, factorial)
	}

	sum := 0
	for _, digit := range factorial {
		sum += digit
	}

	return sum, nil
}

func bigMultiply(x int, n []int) []int {
	carry := 0
	for i := range n {
		product := n[i]*x + carry
		n[i] = product % 10
		carry = product / 10
	}

	for carry > 0 {
		n = append(n, carry%10)
		carry /= 10
	}

	return n
}
