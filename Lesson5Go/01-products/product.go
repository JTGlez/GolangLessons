package main

import "fmt"

type Product struct {
	ID          string
	Name        string
	Price       float64
	Description string
	Category    string
}

type Products []Product

func (p Products) GetAll() {
	fmt.Println("Printing Products...")
	for _, product := range p {
		fmt.Printf("%v\n", product)
	}
	fmt.Println("\n")
}

func (p Product) Save() {
	globalProducts = append(globalProducts, p)
}

func getById(index int) Product {
	return globalProducts[index]
}
