package main

import "fmt"

func main() {

	var ptr *int

	fmt.Println(ptr) //nil

	// & address operator
	// * dereference operator

	var number int = 10
	ptr = &number
	fmt.Printf("Memory address is %p\n", ptr)
	fmt.Printf("Value of number is %d\n", *ptr)

}
