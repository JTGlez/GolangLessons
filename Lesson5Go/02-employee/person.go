package main

import "fmt"

type Person struct {
	ID          string
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       string
	Position string
	Person   Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("ID: %s, Position: %s, Personal ID: %s, Name: %s, Date of Birth: %s \n", e.ID, e.Position, e.Person.ID, e.Person.Name, e.Person.DateOfBirth)
}
