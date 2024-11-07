/* Go is not and objected oriented language. However embedded structs provide a kind
of data-only inheritance that can be useful at times. Keep in mind Go doesn't support classes
or inheritance in the complete sense, embedded structs are just a way to elevate and share
fields between structs */

package main

import "fmt"

type car1 struct {
	make  string
	model string
}

type truck struct {
	car1
	bedSize int
}

func embedded() {

	t := truck{
		bedSize: 1000,
		car1: car1{
			make:  "Ford",
			model: "F150",
		},
	}

	fmt.Println(t.bedSize)
	fmt.Println(t.make)
}
