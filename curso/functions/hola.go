package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

// Mutltipe parameters: When multiple parameters of the same type are passed to a function, the type can be declared only once.
// func concat(s1, s2 string) string {

func userName(firstName, lastName string, age int) string {
	return "Name: " + firstName + " " + lastName + " Age: " + string(age)
}

func main() {
	fmt.Println(concat("Hola", "Mundo"))
	fmt.Println(concat("Go", "Is Fantastic"))
	fmt.Println(userName("John", "Doe", 30))
}
