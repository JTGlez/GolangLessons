package main

import "fmt"

func updateSlice(s []string) {
	s[0] = "Bye"
}

func main() {

	alex := Person{
		FirstName: "Alex",
		LastName:  "Anderson",
		Contact: ContactInfo{
			Email:   "alex@gmail.com",
			ZipCode: 94000,
		},
	}

	alex.Print()
	fmt.Printf("Memory Address of alex: %p\n", &alex)

	alex.UpdateName("Alexander")
	alex.Print()

	fmt.Printf("Memory Address of alex after updating: %p\n", &alex)

	// * Slice is actually a reference type which points to an array dinamically resized

	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice)

	name := "Bill"
	fmt.Println(*&name) // Prints Bill

	// Reference types: they are pointers to another data structure where data is actually stored
	// So, we can change the data safely without worrying about the memory address
	// Examples: slices, maps, channels, pointers, functions

	// Value types: they are the actual data stored in a memory address, so we need to use pointers
	// to change the data stored in this kind of variables
	// Examples: int, float, string, bool, structs

	billPointer := &name
	fmt.Printf("%p\n", billPointer) // Prints Bill

	printPointer(billPointer) // Prints another memory address pointing to name

}

// Actually, go creates a new pointer which points to the previous pointer that we created (billPointer)
func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
