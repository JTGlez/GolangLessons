package main

import (
	"class3go/helpers"
	"errors"
	"fmt"
	"strconv"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	minimum := nums[0]
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

func maxFunc(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	maximum := nums[0]
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

func averageFunc(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func operation(calcType string) (func(nums ...float64) float64, error) {
	switch calcType {
	case minimum:
		return minFunc, nil
	case maximum:
		return maxFunc, nil
	case average:
		return averageFunc, nil
	default:
		return nil, errors.New("calculation type not defined")
	}
}

func main() {
	minFunc, err := operation(minimum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	averageFunc, err := operation(average)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	maxFunc, err := operation(maximum)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Write a sequence of values. Then, press enter to calculate min, max and avg.")

	args := helpers.ReadArgsFromStdin()
	if len(args) == 0 {
		panic("No args provided")
	}

	var values []float64
	for _, arg := range args {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	minValue := minFunc(values...)
	averageValue := averageFunc(values...)
	maxValue := maxFunc(values...)

	fmt.Printf("Valor mínimo: %.2f\n", minValue)
	fmt.Printf("Valor promedio: %.2f\n", averageValue)
	fmt.Printf("Valor máximo: %.2f\n", maxValue)
}
