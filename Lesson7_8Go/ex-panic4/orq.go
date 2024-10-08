package main

import "fmt"

type Orquestator struct {
	handlers map[string]func()
}

func (o *Orquestator) RunHandler(name string) {
	// Panic recovering
	// r stores the data about the generated panic
	// This will stop the panic stack trace
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	hd, ok := o.handlers[name]
	if !ok {
		return
	}

	hd()

}

func HandlerOne() {
	println("handler1")
}

func HandlerTwo() {
	println("handler2")
}

func HandlerThree() {
	panic("handler 3 failed")
}
