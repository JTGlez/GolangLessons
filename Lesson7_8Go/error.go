package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Code    int
}

func (e *CustomError) Error() string {
	return e.Message
}

// err always should be defined as the last return value
func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by 0")
		return
	}
	result = a / b
	return
}

func validateStatusCode(code int) error {
	if code > 399 {
		return errors.New("unexpected http status code")
	}
	return nil
}

func sayHi(name string) (greeting string, err error) {
	if name == "" {
		err = errors.New("no name provided")
		return
	}
	greeting = fmt.Sprintf("Hi, %s.", name)
	return
}
