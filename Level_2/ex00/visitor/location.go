package main

type Location interface {
	accept(Visitor)
}
