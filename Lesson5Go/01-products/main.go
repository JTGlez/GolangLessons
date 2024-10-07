package main

import "fmt"

var globalProducts Products = Products{
	Product{
		ID:          "123",
		Name:        "Xbox",
		Price:       13000,
		Description: "Xbox Series X",
		Category:    "Videogames",
	},
	Product{
		ID:          "124",
		Name:        "PlayStation",
		Price:       14000,
		Description: "PlayStation 5",
		Category:    "Videogames",
	},
}

func main() {

	globalProducts.GetAll()

	addProduct := Product{
		ID:          "125",
		Name:        "Nintendo Switch",
		Price:       10000,
		Description: "Nintendo Switch Console",
		Category:    "Videogames",
	}

	addProduct.Save()

	globalProducts.GetAll()

	fmt.Println(getById(0))

}
