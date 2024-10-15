package main

import "fmt"

// Primitives and structs are concrete types, because we can directly create values of these types.
type EnglishBot struct{}
type SpanishBot struct{}

// Interfaces are abstract types, which can only be used when a struct implements the interface signature.
type Bot interface {
	getGreeting() string
}

// * Interfaces are not generic types!
// * Interfaces are implicit, meaning that we don't have to explicitly say that a struct implements an interface.
// * Interfaces are a contract to help us manage types, but they don't guarantee correctness in the methods.
// * Interfaces are a way to define behavior.

func main() {

	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

// The idea here is that these two implementations of getGreeting() are completely different from each other.
func (EnglishBot) getGreeting() string {
	return "Hello there!"
}

// * If we don't use the receiver, we can remove it from the funciont signature.
func (SpanishBot) getGreeting() string {
	return "Hola!"
}

// In contrast, the printGreeting() function is the same for both EnglishBot and SpanishBot.
/* func printGreeting(eb EnglishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb SpanishBot) {
	fmt.Println(sb.getGreeting())
}
*/
