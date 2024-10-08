package main

import "fmt"

var errorCount int

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrorClientExists || r == ErrorInvalidData {
				fmt.Printf("Error: %v\n", r)
			} else {
				fmt.Printf("Unknown error! %v\n", r)
			}
		}
		fmt.Println("End of execution")
	}()

	/* 	data, err := GetClientData("fulanito.txt")
	   	if err != nil {
	   		fmt.Printf("%s\n", err.Error())
	   		errorCount++
	   	}

	   	fmt.Println("Hola")

	   	fmt.Printf("%+v\n", data)

	   	clientStore.Print()

	   	clientStore.Add(data)

	   	clientStore.Print() */

	data2, err2 := GetClientData("perengano.txt")
	if err2 != nil {
		fmt.Printf("%s\n", err2.Error())
	}

	fmt.Printf("Hola")
	fmt.Printf("%+v\n", data2)

	clientStore.Print()

	clientStore.Add(data2)

	clientStore.Print()
}
