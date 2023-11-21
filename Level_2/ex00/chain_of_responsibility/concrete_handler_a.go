package main

import "fmt"

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
