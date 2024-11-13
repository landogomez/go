/* Formatting Strings review

using standards library's fmt.Sprintf() function. It´s a string interpolation function, similar to JavaScript´s built-in template literals. The %v substring uses the type´s
defauñt formatting, which is iften what you want*/

package main

import "fmt"

func formatting() {
	// Default values
	const name = "Kim"
	const age = 22
	s := fmt.Sprintf("%v is %v years old", name, age) // Sprintf works similarly to Printf, but instead of printing the formatted string to the console, it returns the formatted string.
	fmt.Println(s)

	// Rounding floats
	fmt.Printf("I am %.2f years old\n", 23.321)
}
