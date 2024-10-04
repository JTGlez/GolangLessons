package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

func calcSalary(minutes int, category string) float64 {
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
		fmt.Println("Invalid category")
		return 0
	}

	return baseSalary + bonus
}

func main() {
	fmt.Println("Enter worked minutes and category. Then, press enter...")
	args := helpers.ReadArgsFromStdin()
	if len(args) < 2 {
		fmt.Println("Please provide the worked hours and the employee category.")
		return
	}

	workedHours, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error converting worked hours to integer:", err)
		return
	}

	category := args[1]
	salary := calcSalary(workedHours, category)
	fmt.Printf("Monthly salary: $%.2f\n", salary)
}
