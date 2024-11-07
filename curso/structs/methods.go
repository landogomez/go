/*While go is not object oriented. It does supports methods which can be defined on structs. Methods are
just functions that have a reciever.  A receiver is a special parameter that syntactically goes before
the name of the function*/

package main

import "fmt"

type rect struct {
	height int
	Width  int
}
type authenticationInfo struct {
	username string
	password string
}

func (authi authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authi.username, authi.password)
}

// area has a receiver of type rect (r rect)
func (r rect) area() int {
	return r.height * r.Width
}

func methods() {
	r := rect{height: 10, Width: 5}
	fmt.Println("Area: ", r.area())
	fmt.Println(authenticationInfo{username: "postgres", password: "sasa"}.getBasicAuth())
}
