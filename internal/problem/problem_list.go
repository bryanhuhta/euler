package problem

import (
	"fmt"
)

type problemFunc func(params Params) (interface{}, error)

func Execute(problemID int, params Params) error {
	if problemID < 1 || problemID >= len(problemList) {
		return fmt.Errorf("invalid problem number: must be within [1, %d]", len(problemList)-1)
	}

	val, err := problemList[problemID](params)
	if err != nil {
		return err
	}
	fmt.Println(val)

	return nil
}

func unimplemented(problemNo int) problemFunc {
	return func(params Params) (interface{}, error) {
		err := fmt.Errorf(
			"Problem %d is unimplemented. Try solve it here: https://projecteuler.net/problem=%d",
			problemNo,
			problemNo,
		)
		return nil, err
	}
}

var problemList = []problemFunc{
	nil, // No problem zero
	problem001,
	problem002,
	problem003,
	problem004,
	problem005,
	problem006,
	problem007,
	problem008,
	problem009,
	problem010,
	problem011,
	problem012,
	unimplemented(13),
	unimplemented(14),
	unimplemented(15),
	unimplemented(16),
	unimplemented(17),
	unimplemented(18),
	problem019,
}
