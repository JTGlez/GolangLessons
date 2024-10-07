package main

func main() {
	anEmployee := Employee{

		ID:       "ABC",
		Position: "Software Developer",
		Person: Person{
			ID:          "123",
			Name:        "Yorch",
			DateOfBirth: "28/09/1900",
		},
	}

	anEmployee.PrintEmployee()
}
