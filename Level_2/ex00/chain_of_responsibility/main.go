package main

func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}
	handlerC := &ConcreteHandlerC{}

	handlerA.SetSuccessor(handlerB)
	handlerB.SetSuccessor(handlerC)

	client := &Client{handler: handlerA}

	client.MakeRequest(5)
	client.MakeRequest(15)
	client.MakeRequest(25)
}

// Команда для запуска:
// go run .
