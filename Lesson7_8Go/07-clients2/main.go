package main

import "fmt"

func main() {

	defer fmt.Println("execution finalized")

	content, err := GetClientFile("customers.txt")

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	fmt.Println(content)

}
