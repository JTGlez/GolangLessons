package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

func main() {

	args := helpers.ReadArgsFromStdin()
	if len(args) == 0 {
		panic("No args provided")
	}

	var grades []float64
	for _, arg := range args {
		grade, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			panic(err)
		}
		grades = append(grades, grade)
	}

	average, err := CalculateAverage(grades...)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("Average: %.2f\n", average)
	}

}
