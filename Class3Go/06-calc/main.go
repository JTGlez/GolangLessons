package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

const (
	sum = "+"
	sub = "-"
	mul = "*"
	div = "/"
)

func add(val1, val2 float64) float64 {
	return val1 + val2
}

func subtract(val1, val2 float64) float64 {
	return val1 - val2
}

func multiply(val1, val2 float64) float64 {
	return val1 * val2
}

func divide(val1, val2 float64) float64 {
	if val2 == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return val1 / val2
}

func calculate(operator string) func(val1, val2 float64) float64 {
	switch operator {
	case sum:
		return add
	case sub:
		return subtract
	case mul:
		return multiply
	case div:
		return divide
	default:
		return nil
	}
}

func main() {

	fmt.Println("Enter an operation and two valuesxs")

	args := helpers.ReadArgsFromStdin()
	if len(args) != 3 {
		fmt.Println("Error: You must enter an operator and two values.")
		return
	}

	operator := args[0]
	val1, err1 := strconv.ParseFloat(args[1], 64)
	val2, err2 := strconv.ParseFloat(args[2], 64)

	calcFunc := calculate(operator)(val1, val2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Invalid number format.")
		return
	}

	fmt.Printf("Result of %s: %0.2f\n", operator, calcFunc)

}
