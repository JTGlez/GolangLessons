package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string
	Gender     string
	Age        int
	Profession string
	Likes      Preferences
}

type Preferences struct {
	Foods  string
	Animes string
}

func main() {

	//* Structs -> Related fields together

	p1 := Person{
		Name:       "Jorge",
		Gender:     "Male",
		Age:        26,
		Profession: "Engineer",
	}
	fmt.Println(p1)
	fmt.Printf("My name is %s\n", p1.Name)

	var p3 Person

	p3.Name = "Luis"
	p3.Age = 20
	fmt.Println(p3)
	fmt.Printf("My name is %s\n", p3.Name)

	p4 := Person{
		Name:       "Noreen",
		Gender:     "Female",
		Age:        23,
		Profession: "Doctor",
		Likes: Preferences{
			Foods:  "Chilaquiles",
			Animes: "Doramas",
		},
	}
	fmt.Println(p4)

	p4.Likes.Foods = "Enchiladas"
	fmt.Println(p4)

	p4.Likes = Preferences{
		Foods:  "Tacos",
		Animes: "Jujutsu Kaisen",
	}
	fmt.Println(p4)

	//* Labels -> Annotations after a type declaration

	type Engineer struct {
		Name      string `name:"name vale"`
		Ocupation string `ocupation:"engineer"`
	}

	// * They can be used to establish an encoder from structs -> json representation

	type Subject struct {
		Math    int `json:"math"`
		Biology int `json:"biology"`
	}

	type Student struct {
		FirstName   string  `json:"first_name"`
		LastName    string  `json:"last_name"`
		Scholarship bool    `json:"scholarship,omitempty"` // Won't appear on JSON if no value is given
		Password    string  `json:"-"`                     // Will be ignored while encoding type to a JSON
		Grades      Subject `json:"grades"`
	}

	student := Student{
		FirstName: "Juan",
		LastName:  "Pérez",
		Password:  "12345",
		Grades: Subject{
			Math:    8,
			Biology: 10,
		},
	}

	studentAsJson, _ := json.Marshal(student)
	fmt.Println(string(studentAsJson))

	circleK := Circle{
		radius: 2.5,
	}

	fmt.Printf("Area of circleK is %.2f\n", circleK.area())
	fmt.Printf("Perimeter of circleK is %.2f\n", circleK.perimeter())

	circleK.setRadius(2)

	fmt.Printf("Area of circleK is %.2f\n", circleK.area())
	fmt.Printf("Perimeter of circleK is %.2f\n", circleK.perimeter())

	myCar := Car{
		Engine:        Engine{Horsepower: 250, Type: "V4"},
		Chassis:       Chassis{Material: "Steel"},
		Bodywork:      Bodywork{Color: "Red"},
		NumberOfDoors: 4,
	}

	motorcycle := Motorcycle{
		Engine: Engine{
			Horsepower: 100,
			Type:       "Inline-4",
		},
		Chassis: Chassis{
			Material: "Aluminum",
		},
		Bodywork: Bodywork{
			Color: "Black",
		},
	}

	myCar.Start()
	myCar.DisplayInfo()

	motorcycle.Start()
	motorcycle.DisplayInfo()

	// * Car and Motorcycle, as each one implements the same interface, can be considered as a Vehicle

	var vehicle Vehicle = myCar
	vehicle.Start()
	vehicle.DisplayInfo()

	var bike Vehicle = motorcycle
	bike.Start()
	bike.DisplayInfo()

}
