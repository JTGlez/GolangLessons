package main

import (
	"fmt"
)

type Preferences struct {
	Food, Movie, Sport string
}

type Person struct {
	Name  string
	Age   int
	Likes Preferences
}

func main() {
	p := Person{
		Name: "Tobias",
		Age:  24,
		Likes: Preferences{
			Food:  "Asado",
			Movie: "Matrix",
			Sport: "Swimming",
		},
	}

	fmt.Println(p)

	// Changing movie by reference
	p.Likes.Movie = "Titanic"

	// Taking movie by value, so every reassign won't affect the original struct
	likes := p.Likes
	likes.Movie = "Marvel"
	fmt.Println(p)

	orq := Orchestator{
		Printer: func(ms string) { fmt.Println(ms) },
	}

	orq.Printer("Hello World!")

}
