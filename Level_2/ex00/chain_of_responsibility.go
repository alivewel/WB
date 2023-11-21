package main

import "fmt"

type Handler interface {
	HandleRequest(request int)
	SetSuccessor(successor Handler)
}

type ConcreteHandlerA struct {
	successor Handler
}

func (h *ConcreteHandlerA) HandleRequest(request int) {
	if request <= 10 {
		fmt.Println("ConcreteHandlerA handled the request")
	} else if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}

func (h *ConcreteHandlerA) SetSuccessor(successor Handler) {
	h.successor = successor
}

type ConcreteHandlerB struct {
	successor Handler
}

func (h *ConcreteHandlerB) HandleRequest(request int) {
	if request > 10 && request <= 20 {
		fmt.Println("ConcreteHandlerB handled the request")
	} else if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}

func (h *ConcreteHandlerB) SetSuccessor(successor Handler) {
	h.successor = successor
}

type ConcreteHandlerC struct {
	successor Handler
}

func (h *ConcreteHandlerC) HandleRequest(request int) {
	if request > 20 {
		fmt.Println("ConcreteHandlerC handled the request")
	} else if h.successor != nil {
		h.successor.HandleRequest(request)
	}
}

func (h *ConcreteHandlerC) SetSuccessor(successor Handler) {
	h.successor = successor
}

type Client struct {
	handler Handler
}

func (c *Client) MakeRequest(request int) {
	c.handler.HandleRequest(request)
}

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
