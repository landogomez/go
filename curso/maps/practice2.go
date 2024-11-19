package main

import (
	"errors"
	"fmt"
)

type alumno struct {
	nombre string
	boleta int
}

func getAlumnoMap(nombres []string, boletas []int) (map[int]alumno, error) {
	alumnoMap := make(map[int]alumno)
	if len(nombres) != len(boletas) {
		return nil, errors.New("invalid sizes")
	}

	for i := 0; i < len(nombres); i++ {
		nombre := nombres[i]
		boleta := boletas[i]
		alumnoMap[boleta] = alumno{
			nombre: nombre,
			boleta: boleta,
		}
	}

	return alumnoMap, nil
}

func test2(nombres []string, boletas []int) {
	alumnos, err := getAlumnoMap(nombres, boletas)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, boleta := range boletas {
		fmt.Println("key:", boleta)
		fmt.Println(" - nombre:", alumnos[boleta].nombre)
		fmt.Println(" - boleta:", alumnos[boleta].boleta)
	}
}

func practice2() {
	test2([]string{"alice", "bob", "charlie"}, []int{123, 456, 789})
	test2([]string{"alice", "bob", "charlie"}, []int{123, 456})
	g, _ := getAlumnoMap([]string{"George"}, []int{2022630085})

	value, ok := g[2022630085]
	if ok {
		fmt.Println(value.nombre)
	}
}
