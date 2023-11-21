package main

type Handler interface {
	HandleRequest(request int)
	SetSuccessor(successor Handler)
}
