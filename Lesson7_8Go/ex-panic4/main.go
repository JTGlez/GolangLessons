package main

func main() {
	or := &Orquestator{
		handlers: map[string]func(){
			"handler1": HandlerOne,
			"handler2": HandlerTwo,
			"handler3": HandlerThree,
		},
	}

	// Run handlers
	or.RunHandler("handler1")
	or.RunHandler("handler2")
	or.RunHandler("handler3")
	/* 	or.RunHandler("handler2")
	   	or.RunHandler("handler3") */

}
