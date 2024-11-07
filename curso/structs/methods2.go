// Practicando lo aprendido en el archivo methods.go

package main

import "fmt"

type puerta struct {
	tipo   string
	estado bool
}
type casa struct {
	altura  int
	anchura int
	puerta
}

func (p puerta) abrir() {
	p.estado = true
}

func estadoPuerta(p puerta) string {
	if p.estado {
		return "abierta"
	}
	return "cerrada"
}

func methods2() {
	c := casa{
		altura:  200,
		anchura: 300,
		puerta: puerta{
			tipo: "madera",
		},
	}

	fmt.Printf("La casa tiene una altura de: %v una anchura de: %d y una puerta de: %s\n", c.altura, c.anchura, c.tipo)
	c.abrir()
	fmt.Printf("La puerta esta: %s\n", estadoPuerta(c.puerta))

}
