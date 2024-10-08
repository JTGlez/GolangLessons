package main

import "fmt"

func main() {
	employee := Employee{Hours: 80}
	hourRate := 1000.0

	salary, err := employee.GetSalary(hourRate)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The monthly salary is: %.2f\n", salary)
}
