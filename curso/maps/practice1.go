package main

import "fmt"

func practice1() {
	mapa := map[string]string{
		"George": "5519568270",
		"Manny":  "5519568271",
	}

	// INSERT AN ELEMENT
	mapa["Orlando"] = "5519568272"
	for key, value := range mapa {
		println(key, value)
	}
	fmt.Println("============")

	// GET AN ELEMENT
	phone := mapa["George"]
	println("George Tel. " + phone)
	fmt.Println("============")

	// DELETE AN ELEMENT
	delete(mapa, "Manny")

	for key, value := range mapa {
		println(key, value)
	}

	// fmt.Println(mapa)

	/*
		INSERT AN ELEMENT
		m[key] = value

		GET AN ELEMENT
		value = m[key]

		DELETE AN ELEMENT
		delete(m, key)

		CHECK IF A KEY EXISTS
		value, ok := m[key]
	*/

}
