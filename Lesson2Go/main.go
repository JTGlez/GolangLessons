package main

import (
	"fmt"
)

func main() {

	// * if condition
	var salary float64 = 4000

	if salary <= 3000 {
		fmt.Println("This person must pay taxes")
	} else if salary <= 4000 {
		fmt.Printf("They must pay $%4.2f of their salary \n", (salary/100)*10)
	} else {
		fmt.Printf("They must pay $%4.2f of their salary \n", (salary/100)*10)
	}

	// *switch without expression -> Conditions are controlled in every case
	var age uint8 = 18

	switch {
	case age >= 150:
		fmt.Println("Are you inmortal?")
	case age >= 18:
		fmt.Println("You are an adult.")
	default:
		fmt.Println("You 're not an adult yet.")
	}

	// *switch with top express -> Conditions are given by the expression value
	day := "sunday"

	switch day {
	case "monday", "tuesday":
		fmt.Printf("%s is a bad day :(\n", day)
	default:
		fmt.Printf("%s is a good day :)\n", day)
	}

	// *fallthrough switch -> if the execution enters in a given case, it will execute the next one without checking the condition
	// *expected for the next case (if i is divisible by 15, then it will print Fizz Buzz).
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("Fizz")
			fallthrough
		case i%5 == 0:
			fmt.Print("Buzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		default:
			fmt.Printf("%d", i)
		}
		fmt.Print(" ")
	}

	// * for range -> iterates through an "iterable" structure, like an array or an slice; a dynamic array
	animes := []string{"Konosuba", "Love Live", "Evangelion"}

	for i, anime := range animes {
		fmt.Printf("\nAnime number %d is %s\n", i, anime)
	}

	// * infinite loop
	/* 	sum := 0
	for {
		sum++
	} */

	// * "while" loop -> a for loop with only an stop condition
	sum := 0
	for sum < 10 {
		fmt.Println(sum)
		sum += 1
	}

	// * continue statement -> executes the codeblock given the if condition
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "is odd.")
	}

	// * break -> stops a loop execution (kind of do-while)
	sum2 := 0
	for {
		sum2++
		if sum2 >= 1000 {
			fmt.Println(sum2)
			break
		}
	}

	/* Operators
	* x + y
	* x - y
	* x * y
	* x / y
	* x++
	* y--
	* x == y, x != u, x > y, x < y, <= >=, &&, ||
	* Pointer Address -> &x
	* Pointer access  -> *x
	 */

	// * Arrays -> collection of contiguous elements with a restricted capacity
	var a [1]string
	a[0] = "Hello"
	fmt.Println(a[0])

	// * Slice -> Dinamically sized array
	var s = []bool{true, false}
	fmt.Println(s)

	b := make([]int, 5) // len(b) = 5
	println(cap(b), len(b))
	fmt.Println(b)

	// * Slice Range -> [begin:end] where end is not inclusive
	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes[1:4])
}
