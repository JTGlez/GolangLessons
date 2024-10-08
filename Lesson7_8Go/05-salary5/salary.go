package main

import "errors"

type Employee struct {
	Hours float64
}

func (e Employee) GetSalary(hourRate float64) (salary float64, err error) {

	if e.Hours < 80 {
		err = errors.New("error: the worker cannot have worked less than 80 hours per month")
		return
	}

	if hourRate < 1.0 {
		err = errors.New("error: invalid hour rate")
		return
	}

	salary = e.Hours * hourRate

	if salary <= 150000 {
		return
	}
	salary = salary - salary*.10
	return
}
