// ASSIGNMENT
package main

import "fmt"

func main() {
	sendsSoFar := 35
	const sendstoAdd = 5
	sendsSoFar = incrmenetSends(sendsSoFar, sendstoAdd)
	fmt.Println(sendsSoFar)

	//firstName, lastName := getNames()
	firstName, _ := getNames()          // Ignoring the second return value cause getNames() returns two values
	fmt.Println("Welcome :", firstName) // Aqui me marcaria un error porque no he usado lastName
}

func incrmenetSends(sendssoFar, sendstoAdd int) int {
	return sendssoFar + sendstoAdd
}

// Ignoring return values
func getNames() (string, string) {
	return "John", "Doe"
}

// NAMED RETURN VALUES
func getCords() (x, y int) {
	return // Implicitly returns x, y
}
