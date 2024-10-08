package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Client struct {
	Name  string
	ID    string
	Phone string
	Home  string
}

type Clients []Client

var (
	clientStore = Clients{
		{
			Name:  "fulanito",
			ID:    "123",
			Phone: "5512334532",
			Home:  "monterrey",
		},
	}
	ErrorNonExistentFile = errors.New("customer file doesn't exists on directory")
	ErrorRegisterClient  = errors.New("bad client data")
	ErrorClientExists    = errors.New("client already exists")
	ErrorInvalidData     = errors.New("client data contains zero value")
)

func GetClientData(filename string) (client Client, err error) {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		panic(ErrorNonExistentFile)
	}

	fields := strings.Split(strings.Split(string(dataByte), "\n")[0], ",")
	if len(fields) != 4 {
		panic(ErrorRegisterClient)
	}

	client = Client{
		Name:  fields[0],
		ID:    fields[1],
		Phone: fields[2],
		Home:  fields[3],
	}

	if valid := ValidateClientData(client); !valid {
		panic(ErrorInvalidData)
	}

	return
}

func (cs Clients) Exists(c Client) bool {
	for _, client := range cs {
		if strings.ToLower(client.Name) == strings.ToLower(c.Name) &&
			client.ID == c.ID &&
			client.Phone == c.Phone &&
			strings.ToLower(client.Home) == strings.ToLower(c.Home) {
			return true
		}
	}
	return false
}

func (cs *Clients) Add(c Client) {
	if cs.Exists(c) {
		panic(ErrorClientExists)
	}
	*cs = append(*cs, c)
}

func ValidateClientData(c Client) bool {
	if c.Name == "" || c.ID == "" || c.Phone == "" || c.Home == "" {
		return false
	}
	return true
}

func (c Clients) Print() {
	fmt.Printf("%+v\n", c)
}
