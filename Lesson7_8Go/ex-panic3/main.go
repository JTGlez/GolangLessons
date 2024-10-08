package main

import (
	"fmt"
	"io"
	"os"
)

/* func GetFileData(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(data)

} */

func GetFileData(filename string) (content string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	dataBytes, err := io.ReadAll(f)
	if err != nil {
		return
	}

	content = string(dataBytes)
	return

}

func main() {
	fileName := "test.txt"
	data, err := GetFileData(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	data += " - processed"
	fmt.Println(data)
}
