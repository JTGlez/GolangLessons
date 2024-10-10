package main

import "fmt"

func main() {

	// * Create a map in different ways
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	fmt.Println(colors)

	var colors2 map[string]string

	fmt.Println(colors2)

	colors3 := make(map[string]string)
	fmt.Println(colors3)

	// * Add a key-value pair to a map
	colors["white"] = "#FFFFFF"

	fmt.Println(colors)

	// * Delete a key-value pair from a map
	delete(colors, "red")

	fmt.Println(colors)

	// * Iterate over a map
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Color: %v, Hex: %v\n", color, hex)
	}
}
