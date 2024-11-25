package main

// Declare a num variable of type int with a value of 50.
// Declare a pointer variable called ptr that stores the memory address of num.
// Print the value of num.
// Print the memory address of num.
// Print the value of ptr.
// Print the value of *ptr.
import "fmt"

func ex1() {
	num := 50
	ptr := &num // ptr stores the memory address of num
	fmt.Println(num)
	fmt.Println(&num)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(num) // 100

}

func ex2() {
	x := 50
	var p *int = &x
	var pp **int = &p
	fmt.Println("Valor de la variable x:", x)
	fmt.Println("Valor de p:", *p)
	fmt.Println("Valor de pp:", *pp)
	**pp = 30
	fmt.Println("Nuevo valor de la variable x:", x)
	fmt.Println("Nuevo valor de p:", *p)
	fmt.Println("Nuevo valor de pp:", *pp)
}
