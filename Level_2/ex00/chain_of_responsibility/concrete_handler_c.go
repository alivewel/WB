package main

import "fmt"

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
