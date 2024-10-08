package main

import "os"

func GetClientFile(filename string) (data string, err error) {

	dataByte, err := os.ReadFile(filename)

	if err != nil {
		return
	}

	data = string(dataByte)
	return

}
