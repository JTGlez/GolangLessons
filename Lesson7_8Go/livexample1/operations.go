package main

import (
	"errors"
	"fmt"
)

type MathError struct {
	a, b int
}

func (m MathError) Error() string {
	return fmt.Sprintf("Math error with values a: %d, b: %d", m.a, m.b)
}

func AddNum(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) {

	if a == 1000 {
		err := fmt.Errorf("something extra boring. %w", &MathError{1000, 1000})
		return 0, err
	}

	if b == 0 {
		if a == 0 {
			err := fmt.Errorf("0, 0. %w", ErrDivisionIndeterminated)
			err = fmt.Errorf("Boring detail, 1. %w", err)
			err = fmt.Errorf("Boring detail, 2. %w", err)
			return 0, err
		} else if a > 0 {
			return 0, ErrDivisionUndefinedPlusInf
		} else {
			return 0, ErrDivisionUndefinedMinusInf
		}
	}
	return a / b, nil

}

// Good practice: declare our errors in this way
// So every time we compare errors, they will point to the same memory address
// In this situacion, using global variables is a good approach
var (
	ErrDivisionIndeterminated    = errors.New("indeterminated")
	ErrDivisionUndefinedPlusInf  = errors.New("undefined, limit tends to +inf")
	ErrDivisionUndefinedMinusInf = errors.New("undefined, limit tends to -inf")
)
