package problem

import (
	"fmt"
	"math"
)

func problem004(params Params) (interface{}, error) {
	err := params.assertKeysExist("digits")
	if err != nil {
		return nil, err
	}

	digits, err := params.getInt("digits")
	if err != nil {
		return nil, err
	}
	digits -= 1

	interval := int(math.Pow10(digits)) // 10^n
	max := int(math.Pow10(digits + 1))  // 10^(n+1)
	min := max - interval

	for i := max - 1; i >= min; i-- {
		for j := i; j >= min; j-- {
			if isPalindrome(i * j) {
				return i * j, nil
			}
		}
	}

	return 0, nil
}

func isPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	for start, end := 0, len(s)-1; start < end; start, end = start+1, end-1 {
		if s[start] != s[end] {
			return false
		}
	}
	return true
}
