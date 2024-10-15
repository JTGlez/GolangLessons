package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	/* A Reader is used to write something in a byte slice and use it inside of the program.
	 * On the other hand, a Writer is used to take a byte slice and send it somewhere else outside of the program.
	 */

	// io.Copy() is a function that takes two arguments: a Writer and a Reader.
	// As a result, it reads the response body and writes it to the standard output.
	io.Copy(os.Stdout, resp.Body)

	/* 	bs := make([]byte, 99999)
	   	// Response implements the Reader interface, so we can use the Read() method.
	   	resp.Body.Read(bs)
	   	// HTML of the Google homepage.
	   	fmt.Println(string(bs)) */

}
