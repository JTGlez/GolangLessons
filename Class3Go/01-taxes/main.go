package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

func calculateTaxes(salary float64) (taxe float64) {
	if salary < 50000 {
		taxe = 0
	} else if salary <= 15000 {
		taxe = salary * 0.17
	} else {
		taxe = salary * 0.27
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
		fmt.Printf("Para un salario de $%.2f, el impuesto es $%.2f\n", salary, taxes)
	}
}
