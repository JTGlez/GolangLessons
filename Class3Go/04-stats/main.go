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

func operation(calcType string) (func(nums ...int) int, error) {
	switch calcType {
	case minimum:
		return func(nums ...int) int {
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
		}, nil
	case maximum:
		return func(nums ...int) int {
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
		}, nil
	case average:
		return func(nums ...int) int {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, num := range nums {
				sum += num
			}
			return sum / len(nums)
		}, nil
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

	args := helpers.ReadArgsFromStdin()
	if len(args) == 0 {
		panic("No args provided")
	}

	var values []int
	for _, arg := range args {
		value, err := strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		values = append(values, value)
	}

	minValue := minFunc(values...)
	averageValue := averageFunc(values...)
	maxValue := maxFunc(values...)

	fmt.Printf("Valor mínimo: %d\n", minValue)
	fmt.Printf("Valor promedio: %d\n", averageValue)
	fmt.Printf("Valor máximo: %d\n", maxValue)

}
