package main

type GeometricShape interface {
	Area()
}

type Rectangle struct {
	height int
	base   int
}

func (r Rectangle) Area() int {
	return r.base * r.height
}
