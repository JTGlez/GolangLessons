package main

import (
	"class3go/helpers"
	"errors"
	"fmt"
	"strconv"
)

func calculateAverage(grades ...float64) (average float64, err error) {

	if len(grades) == 0 {
		return 0, errors.New("no grades provided")
	}

	sum := 0.0
	for _, grade := range grades {
		if grade < 0 {
			return 0, errors.New("invalid grade")
		}
		sum += grade
	}
	average = sum / float64(len(grades))
	return
}

func main() {

	args := helpers.ReadArgsFromStdin()
	if len(args) == 0 {
		panic("No args provided")
	}

	var grades []float64
	for _, arg := range args {
		salary, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			panic(err)
		}
		grades = append(grades, salary)
	}

	average, err := calculateAverage(grades...)

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Printf("Average: %.2f\n", average)
	}

}
