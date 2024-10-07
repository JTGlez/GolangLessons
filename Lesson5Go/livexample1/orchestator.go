package main

type Orchestator struct {
	Printer func(ms string)
}

func (o Orchestator) Print(msg string) {
	o.Printer("prefix" + msg)
}
