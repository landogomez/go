package main

import "fmt"

/*
Tarea:

Define Structures, Interfaces and Methods:

Built a structure Emplyee with fields Name, Role and Salary.
Define an interface Promotable with a Promote() method.
Implement Promote() on Employee, Promote() should update the Employee's role and it's salary.
Implement a Display() function that shows the Employee info.

Implement a GeneralPromotion() func that accepts a Promotable and promotes it, using the interface

*/

type employee struct {
	name   string
	role   string
	salary float64
}

type promotable interface {
	promote(r string)
}

func (e *employee) promote(r string) {
	e.role = r
	e.salary *= 2.1
}

func (e employee) Display() {
	fmt.Println("Name: " + e.name)
	fmt.Println("Role: " + e.role)
	fmt.Println("Salary :", e.salary)
}

func GeneralPromotion(e promotable, r string) {
	e.promote(r)
}

func ex4() {
	e := employee{
		name:   "George",
		role:   "Jr Dev",
		salary: 150.000,
	}

	e.Display()
	GeneralPromotion(&e, "Sr Dev")
	e.Display()
	e.promote("Scrum Master")
	e.Display()
}
