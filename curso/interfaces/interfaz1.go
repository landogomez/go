package main

import "fmt"

type animal struct {
	name  string
	color string
}

type dog struct {
	animal
	race string
}
type cat struct {
	animal
	age int
}
type bird struct {
	animal
	weight float64
}

func (c cat) getSound() string {
	return "Meow"
}

func (d dog) getSound() string {
	return "Woof"
}

type sound interface {
	getSound() string
}

func makeSound(s sound) {
	fmt.Println(s.getSound())
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func interfaz1() {
	makeSound(dog{
		animal: animal{
			name:  "Firulais",
			color: "Black",
		},
	})
	makeSound(cat{
		animal: animal{
			name:  "Garfield",
			color: "Orange",
		},
	})

	// Type assertions in GO
	// You can use type assertions to check if an interface value holds a specific type.
	//c, ok := s.(circle)  This is a type assertion. It checks if the interface s holds a circle type. If it does, it assigns the value to c and ok is true. If it doesn't, ok is false.

}

// Type assertions in GO
// You can use type assertions to check if an interface value holds a specific type.
