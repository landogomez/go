package main

// DEFER
// The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before
// its enclosing function returns. This is useful for closing files, releasing resources, and other cleanup tasks.

import (
	"fmt"
	"os"
)

func defer1() {
	// Abrimos un archivo
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Aseguramos que el archivo se cierre al final de la función main
	defer file.Close()

	// Realizamos operaciones con el archivo
	// ...

	fmt.Println("Archivo abierto y operaciones realizadas")
}

func defer2() {
	defer fmt.Println("Primero en diferir, último en ejecutarse")
	defer fmt.Println("Segundo en diferir, penúltimo en ejecutarse")
	defer fmt.Println("Tercero en diferir, primero en ejecutarse")

	fmt.Println("Función main ejecutándose")
}
