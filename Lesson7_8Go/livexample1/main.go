package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := Div(1000, 0)

	if err != nil {

		// * How to know which error is this
		// Don't do this! if err == errors.New("intedeterminated") we 're creating another pointer
		// so it's a false comparison

		//* Type assertion of an error
		// err = errors.Unwrap(err)
		var e *MathError // Will store a pointer to MathError

		// Is checks by message and type

		// As checks only by type

		// Will recursively ask: is err equal to e? then, it will return the error detail
		if errors.As(err, &e) {
			fmt.Println(err)
			return
		}

		_, ok := err.(*MathError)
		if ok {
			fmt.Println("yes, im a math error")
		}

		// Is allow us to unwrap an error contained in a error chain in a recursive way
		if errors.Is(err, ErrDivisionIndeterminated) {
			fmt.Println("code 1: exit")
		} else if err == ErrDivisionUndefinedPlusInf {
			fmt.Println("code 2: exit")
		} else if err == ErrDivisionUndefinedMinusInf {
			fmt.Println("code 3: exit")
		} else {
			fmt.Println("unknown code: exit")
		}

		/* if err == error(MathError{0, 0}) {
			fmt.Println("code 1: exit")
		} else {
			fmt.Println("unknown code: exit")
		} */

		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)

}
