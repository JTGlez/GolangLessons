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

func (b Book) ShowInfo() string {
	return fmt.Sprintf(
		"%s was written by %s and has %d pages - total publishers: %d\n",
		b.Title,
		b.Author,
		b.Pages,
		b.TotalFounders(),
	)
}

type Movie struct {
	Title      string
	Director   string
	IMDBRating float64
}

func main() {

	book := Book{
		Title:  "Harry Potter",
		Author: "J.K. Rowling",
		Pages:  254,
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

	fmt.Printf(book.ShowInfo())

}
