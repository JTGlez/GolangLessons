package main

func CalculateTaxes(salary float64) float64 {
	if salary < 50000 {
		return 0
	} else if salary <= 150000 {
		return salary * 0.17
	} else {
		return salary * 0.27
	}
}
