package main

/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:


Name: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.

Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

func main() {

	stud := Student{
		FirstName: "Leonardo",
		LastName:  "Cruz",
		DNI:       "1234556",
		Date:      "28/09/1990",
	}

	stud.Detalle()

}
