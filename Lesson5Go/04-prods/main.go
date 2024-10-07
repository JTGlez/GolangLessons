package main

import (
	"fmt"
)

func main() {
	smallProduct := newProduct(smallType, 100.0)
	mediumProduct := newProduct(mediumType, 100.0)
	largeProduct := newProduct(largeType, 100.0)

	fmt.Printf("Small Product Price: $%.2f\n", smallProduct.Price())
	fmt.Printf("Medium Product Price: $%.2f\n", mediumProduct.Price())
	fmt.Printf("Large Product Price: $%.2f\n", largeProduct.Price())
}
