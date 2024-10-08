package main

import (
	"errors"
	"fmt"
)

func main() {

	salary := 9000
	var customError error
	targetError := errors.New("error: salary is less than 10000")

	if salary <= 10000 {
		customError = targetError
	}

	if errors.Is(customError, targetError) {
		panic(customError)
	}

	fmt.Println("Salary is above the minimum threshold")

}
