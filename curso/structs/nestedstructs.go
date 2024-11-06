package main

type messageToSendd struct {
	message   string
	sender    user
	recipient user
}
type user struct {
	name   string
	number int
}

func canSendMessage(mensaje messageToSendd) bool {
	if mensaje.sender.name == "" {
		return false
	}
	if mensaje.sender.number == 0 {
		return false
	}
	if mensaje.recipient.name == "" {
		return false
	}
	if mensaje.recipient.number == 0 {
		return false
	}

	return true
}

func nested() {

	x := messageToSendd{} // x is initialized to its zero value
	x.sender.name = "John"
	x.sender.number = 1234567890
	x.recipient.name = ""
	x.recipient.number = 9876543210

	if canSendMessage(x) {
		x.message = "Hello World"
		print("Sending message: ", x.message, " to: ", x.recipient.name)
	} else {
		print("Cannot send message")
	}
}
