package main

import "errors"

func CalculateAverage(grades ...float64) (average float64, err error) {

	if len(grades) == 0 {
		return 0, errors.New("no grades provided")
	}

	sum := 0.0
	for _, grade := range grades {
		if grade < 0 || grade > 10 {
			return 0, errors.New("invalid grade")
		}
		sum += grade
	}
	average = sum / float64(len(grades))
	return
}
