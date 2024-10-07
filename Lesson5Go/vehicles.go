package main

import "fmt"

// * Composition is used as a form of "inheritance" between structs.
// * However, it does not comply with the Liskov Substitution Principle, }
// * since where an "embedded" struct is expected, you cannot substitute it with the "parent" struct.

// * Example: you have a "struct" vehicle which is composed by 3 elements: engine, chassis and bodywork
// * Then, we expect to have different vehicles, like bikes, trucks... (they should behave differently each other)
// * In Go, to achieve such behavior we need to use Interfaces to define common behaviors and composition to reuse functionality

// * So, let´s write the struct for each element of a vehicle

type Engine struct {
	Horsepower int
	Type       string
}

type Chassis struct {
	Material string
}

type Bodywork struct {
	Color string
}

// * Now, let´s write the vehicle interface, which represents a common "behavior" or contract each vehicle could implement

type Vehicle interface {
	Start()
	DisplayInfo()
}

// * Now, we can create "Vehicle" structs, which are composed by the 3 first structs. Then, we 'll implement the interface signature

type Car struct {
	Engine
	Chassis
	Bodywork
	NumberOfDoors int
}

func (c Car) Start() {
	fmt.Println("The car is starting with a roar!")
}

func (c Car) DisplayInfo() {
	fmt.Printf("This car has %d doors, a %s body, a %s chassis, and a %d horsepower %s engine.\n", c.NumberOfDoors, c.Bodywork.Color, c.Chassis.Material, c.Engine.Horsepower, c.Engine.Type)
}

type Motorcycle struct {
	Engine
	Chassis
	Bodywork
}

func (m Motorcycle) Start() {
	fmt.Println("The motorcycle is starting with a vroom!")
}

func (m Motorcycle) DisplayInfo() {
	fmt.Printf("This motorcycle has a %s body, a %s chassis, and a %d horsepower %s engine.\n", m.Bodywork.Color, m.Chassis.Material, m.Engine.Horsepower, m.Engine.Type)
}
