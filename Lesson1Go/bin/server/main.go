package main

// Import by complete path module using Mayus on exports
import (
	"github.com/JTGlez/Lesson1Go/lib/calculator"
)

var config int = 0

func main() {

	println(calculator.Sum)

	// Types
	// - int
	// - float
	// - string
	// - bool
	// - rune

	var firstName string = "Lalo"
	var lastName string

	println(firstName, lastName)

	// Implicit declaration
	var age = 115
	isValid := false
	println(age, isValid)

}
