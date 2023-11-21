package main

import "fmt"

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
