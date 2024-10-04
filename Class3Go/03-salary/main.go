package main

import (
	"class3go/helpers"
	"fmt"
	"strconv"
)

func calcSalary(minutes int, category string) float64 {
	workedHours := float64(minutes) / 60.0
	var baseSalary, bonus float64

	switch category {
	case "C":
		baseSalary = workedHours * 1000
	case "B":
		baseSalary = workedHours * 1500
		bonus = baseSalary * 0.20
	case "A":
		baseSalary = workedHours * 3000
		bonus = baseSalary * 0.50
	default:
		fmt.Println("Invalid category")
		return 0
	}

	return baseSalary + bonus

}

func main() {
	fmt.Println("Enter worked hours and category. Then, press enter...")
	args := helpers.ReadArgsFromStdin()
	if len(args) < 2 {
		fmt.Println("Por favor, proporciona las horas trabajadas y la categoría del empleado.")
		return
	}

	horasTrabajadas, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error al convertir las horas trabajadas a entero:", err)
		return
	}

	categoria := args[1]
	salario := calcSalary(horasTrabajadas, categoria)
	fmt.Printf("Salario mensual: $%.2f\n", salario)
}
