package main

import (
	"errors"
	"fmt"
)

type protocolo struct {
	id     int
	nombre string
	alumno string
}

func getMapProtocolo(ids []int, nombres []string, alumnos []string) (map[int]protocolo, error) {
	protocoloMap := make(map[int]protocolo)
	if len(ids) != len(nombres) || len(ids) != len(alumnos) {
		return nil, errors.New("Invalid size")
	}

	for i := 0; i < len(ids); i++ {
		nombre := nombres[i]
		id := ids[i]
		alumno := alumnos[i]
		protocoloMap[id] = protocolo{
			id:     id,
			nombre: nombre,
			alumno: alumno,
		}
	}

	return protocoloMap, nil
}

// Create a map with yout contacts, it must have name and phoneNumber, for each contact, print name and phone
func practice3() {
	contacts := make(map[string][]int)
	contacts["George"] = []int{5519568270, 56781016}
	contacts["Derp"] = []int{5521010225, 56781016}

	for number, phone := range contacts {
		fmt.Printf("Name: %s\n ", number)
		for i, num := range phone {
			fmt.Printf("Telefono %d: %d\n", i+1, num)
		}
	}
}
