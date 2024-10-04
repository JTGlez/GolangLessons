package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

func calculateTaxes(salary float64) (tax float64) {
	if salary < 50000 {
		tax = 0
	} else if salary <= 150000 {
		tax = salary * 0.17
	} else {
		tax = salary * 0.27
	}
	return
}

func main() {
	args := helpers.ReadArgsFromStdin()
	if len(args) == 0 {
		panic("No arguments provided")
	}

	var salaries []float64
	for _, arg := range args {
		salary, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			panic(err)
		}
		salaries = append(salaries, salary)
	}

	for _, salary := range salaries {
		taxes := calculateTaxes(salary)
		fmt.Printf("For a salary of $%.2f, the tax is $%.2f\n", salary, taxes)
	}
}
