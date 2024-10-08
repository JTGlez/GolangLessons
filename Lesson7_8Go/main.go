package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {

	// Custom error creation
	var ErrorCustom error = &CustomError{"Error Message", 1}
	fmt.Println(ErrorCustom.Error())

	// Use stdlib to create error datatypes
	err := errors.New("Error message with package errors")
	fmt.Println(err)

	// Use fmt to create a formatted error message
	err2 := fmt.Errorf("Error message: %s", "Error with fmt")
	fmt.Println(err2)

	// * Identifying errors -> using == to compare errors or Is() and As()
	// * Errors are defined by a hierarchy using wrappers

	errBase := errors.New("Error base")
	errWrap := fmt.Errorf("Error wrap: %w", errBase)

	// Is (err, target) -> searchs in errWrap/target if errBase/err is defined.
	fmt.Println(errors.Is(errWrap, errBase))

	errBase2 := &CustomError{"Error custom base", 2}
	errWrap2 := fmt.Errorf("Error wrap: %w", errBase2)

	// Example, checking an error while opening a nonexistent file
	// We re checking if the error we got in os.Open is defined on fs.ErrNotExist
	_, errFile := os.Open("i_dont_exist.txt")
	var notExist error = fs.ErrNotExist
	if errors.Is(errFile, notExist) {
		fmt.Println("The file doesn't exist!")
	}

	// (err error, target)
	// As -> compares if an error is equal to our search target, or if target points to the error sent in the first param
	var err3 *CustomError
	if errors.As(errWrap2, &err3) {
		fmt.Println(err3.Code)
	}

	// Another example -> we re comparing if PathError points to the same error that errFile2 references
	_, errFile2 := os.Open("i_dont_exist.txt")
	var pathError *fs.PathError
	if errors.As(errFile2, &pathError) {
		fmt.Println("Path error", pathError.Path)
	}

	// * Error unwrapping to check if it has another error nested
	error1 := fmt.Errorf("first error")
	error2 := fmt.Errorf("second error: %w", error1)
	errDe := errors.Unwrap(error2)
	if errDe == error1 {
		fmt.Println("We got an error1! ", error1)
	}

	// * Functions should use multiple return values to handle potencial failures in function execs
	statusCode := 200
	if err := validateStatusCode(statusCode); err != nil {
		fmt.Printf("http request failed with error: %v\n", err)
		return
	}
	fmt.Println("http code 404!")

	result, err := divide(10, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)

	greeting, errG := sayHi("Cortana")

	if errG != nil {
		fmt.Println(err)
	}

	fmt.Println(greeting)

}
