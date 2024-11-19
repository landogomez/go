package main

import "fmt"

func practice1() {
	carros := [3]string{"BMW", "Merceded", "Audio"}
	carrosSlice := carros[:]
	carrosSlice = append(carrosSlice, "Ferrari")

	for _, carro := range carros {
		fmt.Println(carro)
	}

	array := [3]int{1, 2, 3}
	slice := array[:]
	slice = append(slice, 4)

	for _, valor := range slice {
		fmt.Println(valor)
	}

	transformers := make([]string, 5)
	transformers[0] = "Optimus Prime"
	transformers[1] = "Megatron"
	transformers[2] = "Bumblebee"
	transformers[3] = "Starscream"
	transformers[4] = "Ironhide"
	transformers = append(transformers, "Soundwave", "Ratchet")
	fmt.Println(len(transformers))                      // len equals 7 even though we first created a slice with a length of 5
	transformers = append(transformers, carrosSlice...) // ... is the spread operator.

	for _, transformer := range transformers {
		fmt.Println(transformer)
	}

	transformers = append(transformers[:7], transformers[len(transformers):]...) // Removing multiple elements from a slice
	fmt.Println(transformers)

	carritos := append(carrosSlice, "Lamborghini")
	for _, carrito := range carritos {
		fmt.Println(carrito)
	}

}
