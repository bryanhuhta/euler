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
}
