package main

import "fmt"

type Publisher struct {
	Name     string
	Founders []string
}

func (p Publisher) TotalFounders() int {
	return len(p.Founders)
}

type Book struct {
	Title  string
	Author string
	Pages  int
	Publisher
}

type Movie struct {
	Title      string
	Director   string
	IMDBRating float64
}

func main() {

	book := Book{
		Title: "Harry Potter",
		Publisher: Publisher{
			Name: "Trillas",
			Founders: []string{
				"Fulanito",
				"Perengano",
			},
		},
	}

	fmt.Println("Number of founders:", book.TotalFounders())
	fmt.Println(book.Pages)

}
