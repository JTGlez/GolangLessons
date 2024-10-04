package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

type Animal struct {
	name string
	kg   float64
}

var (
	dog     = Animal{name: "Dog", kg: 10}
	cat     = Animal{name: "Cat", kg: 5}
	hamster = Animal{name: "Hamster", kg: 0.25}
	spider  = Animal{name: "Spider", kg: 0.15}
)

var animals = map[string]Animal{
	"Dog":     dog,
	"Cat":     cat,
	"Hamster": hamster,
	"Spider":  spider,
}

func (d Animal) animal() func(count int) float64 {
	return func(count int) (amount float64) {
		amount = d.kg * float64(count)
		return
	}
}

func main() {

	fmt.Println("Introduce the animal and the number of animals")
	args := helpers.ReadArgsFromStdin()

	if len(args) != 2 {
		panic("Wrong argument count provided")
	}

	animalName := args[0]
	count, err := strconv.Atoi(args[1])

	if err != nil || count <= 0 {
		panic("Invalid number format")
	}

	animal, exists := animals[animalName]
	if !exists {
		panic("Animal not found")
	}

	calcFood := animal.animal()
	totalFood := calcFood(count)
	fmt.Printf("Total food for %d %s(s): %.2f kg\n", count, animal.name, totalFood)

}
