package problem

import (
	"regexp"
	"strconv"
	"strings"
)

func problem011(params Params) (interface{}, error) {
	err := params.assertKeysExist("file")
	if err != nil {
		return nil, err
	}

	text, err := params.getFileText("file")
	if err != nil {
		return nil, err
	}

	rows := parseInts(text)

	// Start at the second from bottom row
	//   For each value, sum each of the two possible child values
	//   Replace the value with the greater sum
	// Move to the next parent row, repeat until at the root
	// The largest path is now at the root
	for i := len(rows) - 2; i >= 0; i-- {
		row := rows[i]

		for j, num := range row {
			leftSum := rows[i+1][j] + num
			rightSum := rows[i+1][j+1] + num

			if leftSum > rightSum {
				row[j] = leftSum
			} else {
				row[j] = rightSum
			}
		}
	}

	return rows[0][0], nil
}

func parseInts(text string) [][]int {
	regex := regexp.MustCompile(`\d*`)

	ints := make([][]int, 0)
	for _, line := range strings.Split(text, "\n") {
		matches := regex.FindAllString(line, -1)

		digits := make([]int, 0)
		for _, match := range matches {
			n, err := strconv.Atoi(match)
			if err != nil {
				continue
			}

			digits = append(digits, n)
		}

		if len(digits) > 0 {
			ints = append(ints, digits)
		}
	}

	return ints
}
