package main

import "fmt"

// Pointer Receivers

/* A receiver type on a method can be a pointer.
Methods with pointer receivers can modify the value to which the reciever points.
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
*/

type car struct {
	color string
}

func (c *car) paint(color string) { // pointer receiver
	c.color = color
}

func pointerReceivers() {
	c := car{
		color: "red",
	}
	fmt.Println(c.color) // red
	c.paint("blue")
	fmt.Println(c.color) // blue
}

// paint() acts like a set method, changing the value of the receiver to the new color.

// ASSIGNMENT

/* setMessages sets the message field of the given email structure, and the new value persists outside the scope of the function. */

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test2(e *email, newMessage string) {
	fmt.Println("--- Before ---")
	e.print()
	fmt.Println("---end Before ---")
	e.setMessage(newMessage)
	fmt.Println("--- After ---")
	e.print()
	fmt.Println("---end After ---")
	fmt.Println("====================================")

}

func (e email) print() {
	fmt.Println("Message:", e.message)
	fmt.Println("From:", e.fromAddress)
	fmt.Println("To:", e.toAddress)
}

func assignment3() {
	e := email{
		message:     "Hello",
		fromAddress: "Yo",
		toAddress:   "Tu",
	}

	test2(&e, "Goodbye")
}
