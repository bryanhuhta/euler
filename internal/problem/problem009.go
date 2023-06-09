package problem

func problem009(params Params) (interface{}, error) {
	err := params.assertKeysExist("max")
	if err != nil {
		return nil, err
	}

	max, err := params.getInt("max")
	if err != nil {
		return nil, err
	}

	product := 0
	triples := pythagoreanTriples(max)

	for _, triple := range triples {
		a, b, c := triple[0], triple[1], triple[2]
		if a+b+c == 1000 {
			product = a * b * c
			break
		}
	}

	return product, nil
}

func pythagoreanTriples(max int) [][]int {
	triples := make([][]int, 0)

	for a := 1; a < max; a++ {
		for b := a; b < max; b++ {
			for c := b; c < max; c++ {
				if a*a+b*b == c*c {
					triples = append(triples, []int{a, b, c})
				}
			}
		}
	}

	return triples
}
