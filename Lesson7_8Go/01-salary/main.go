package main

import "fmt"

func main() {

	salary := 100000
	// Pointer -> Implicit reference to SalaryError
	var CustomError error = &SalaryError{
		Message: "Error: the salary entered does not reach the taxable minimum",
		Code:    1,
	}

	if salary < 150000 {
		panic(CustomError)
	}

	fmt.Println("Must pay tax")

}
