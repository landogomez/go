/* An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be reference
elsewhere */

package main

import "fmt"

type car struct {
	Make   string
	Model  string
	Height int
	Width  int
	Wheel  struct {
		Radius   int
		Material string
	}
}

// When to use?
// When you need a struct for a single use case and you don't need to reference it elsewhere

func anonymous() {
	c := car{
		Make:   "BMW",
		Model:  "M4",
		Height: 1400,
		Width:  2000,
		Wheel: struct {
			Radius   int
			Material string
		}{
			Radius:   18,
			Material: "Alloy",
		},
	}

	// Accessing the fields
	fmt.Println(c.Wheel.Radius) // 18
}
