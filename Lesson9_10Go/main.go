package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Third argument is the permission: only applicable when creating a file
	r, err := os.OpenFile("something.txt", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer r.Close()

	s := []byte{}

	for {
		// buf is actually a slice of bytes, which also points to an array
		buf := make([]byte, 8)
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("EOF: %v\n", err)
				break
			}
			fmt.Printf("Unknown error: %v", err)
			return
		}
		// Define n to avoid reading the buffer's empty spaces and waste memory
		s = append(s, buf[:n]...)
	}

	fmt.Println(string(s))

	/* r := strings.NewReader("hello, how are you?")
	fmt.Printf("%+v", r)

	buf := make([]byte, 4)
	n, err := r.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf[:n])
	fmt.Println(string(buf[:n]))

	n, err = r.Read(buf)
	fmt.Println(string(buf[:n]))

	n, err = r.Read(buf)
	fmt.Println(string(buf[:n]))

	n, err = r.Read(buf)
	fmt.Println(string(buf[:n]))

	n, err = r.Read(buf)
	fmt.Println(string(buf[:n]))

	n, err = r.Read(buf)
	fmt.Println(string(buf[:n])) */
}
