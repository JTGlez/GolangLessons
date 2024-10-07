package main

import "math"

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Changing radius by reference
func (c *Circle) setRadius(newRadius float64) {
	c.radius = newRadius
}
