package main

import "fmt"

type ContactInfo struct {
	Email   string
	ZipCode int
}

type Person struct {
	FirstName string
	LastName  string
	Contact   ContactInfo
}

func (p Person) Print() {
	fmt.Printf("%+v\n", p)
}

func (p *Person) UpdateName(newFirstName string) {
	(*p).FirstName = newFirstName
}
