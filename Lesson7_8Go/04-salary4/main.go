package main

import (
	"errors"
	"fmt"
)

func main() {

	salary := 9000
	var customError error
	targetError := fmt.Errorf("error: the minimum taxable amount is 150,000 and the salary entered is: %d", salary)

	if salary <= 10000 {
		customError = targetError
	}

	if errors.Is(customError, targetError) {
		panic(customError)
	}

	fmt.Println("Salary is above the minimum threshold")

}
