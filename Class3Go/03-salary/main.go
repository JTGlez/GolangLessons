package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter worked minutes and category. Then, press enter...")
	args := helpers.ReadArgsFromStdin()
	if len(args) < 2 {
		fmt.Println("Please provide the worked hours and the employee category.")
		return
	}

	workedHours, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error converting worked hours to integer:", err)
		return
	}

	category := args[1]
	salary, _ := CalcSalary(workedHours, category)
	fmt.Printf("Monthly salary: $%.2f\n", salary)
}
