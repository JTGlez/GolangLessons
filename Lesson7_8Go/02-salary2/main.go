package main

import "errors"

func main() {

	salary := 90000
	var CustomError error

	targetError := &SalaryError{
		Message: "Error: salary is less than 10000",
		Code:    1,
	}

	otherError := errors.New("another dumb error")

	errorWrap := ErrorWrapper{
		TargetError: targetError,
		OtherError:  otherError,
	}

	if salary <= 100000 {
		CustomError = targetError
	}

	if errors.Is(CustomError, errorWrap.TargetError) {
		panic(CustomError)
	}

}
