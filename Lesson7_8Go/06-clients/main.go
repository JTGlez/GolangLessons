package main

import (
	"fmt"
)

func main() {

	// Executes them as a LIFO Stack
	defer fmt.Println("soy el seguns")
	defer fmt.Println("primero salgo")

	content, err := GetClientFile("customers.txt")

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	fmt.Println(content)

}
