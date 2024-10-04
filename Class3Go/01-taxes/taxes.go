package main

import "errors"

func CalculateTaxes(salary float64) (float64, error) {
	if salary <= 0 {
		return 0, errors.New("invalid salary")
	} else if salary < 50000 {
		return 0, nil
	} else if salary <= 150000 {
		return salary * 0.17, nil
	} else {
		return salary * 0.27, nil
	}
}
