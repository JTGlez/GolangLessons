package main

import "fmt"

func CalcSalary(minutes int, category string) (float64, error) {
	if minutes < 0 {
		return 0, fmt.Errorf("minutes cannot be negative")
	}

	workedHours := float64(minutes) / 60.0
	var baseSalary, bonus float64

	switch category {
	case "C":
		baseSalary = workedHours * 1000
	case "B":
		baseSalary = workedHours * 1500
		bonus = baseSalary * 0.20
	case "A":
		baseSalary = workedHours * 3000
		bonus = baseSalary * 0.50
	default:
		return 0, fmt.Errorf("invalid category")
	}

	return baseSalary + bonus, nil
}
