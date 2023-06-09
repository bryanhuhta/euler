package problem

import (
	"time"
)

func problem019(params Params) (interface{}, error) {
	err := params.assertKeysExist("end")
	if err != nil {
		return nil, err
	}

	endYear, err := params.getInt("end")
	if err != nil {
		return nil, err
	}

	sundays := 0
	for year := 1900; year < endYear; year++ {
		firstSundays := 0
		for month := time.January; month <= time.December; month++ {
			date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
			if date.Weekday() == time.Sunday {
				firstSundays++
			}
		}

		if year != 1900 {
			sundays += firstSundays
		}
	}

	return sundays, nil
}
