package main

import (
	"fmt"
)

type Receiver struct{}

func (c Receiver) action() {
	fmt.Println("Recevier: execute action")
}

type Command interface {
	execute()
	undo()
}

type ConcreteCommand struct {
	executed bool
	r        Receiver
}

func (cc ConcreteCommand) execute() {
	if !cc.executed {
		cc.r.action()
		cc.executed = true
	}
}

func (cc ConcreteCommand) undo() {
	if cc.executed {
		fmt.Println("ConcreteCommand: undo action")
		cc.executed = false
	}
}

type Invoker struct {
	c Command
}

func (i Invoker) setComand(com Command) {
	i.c = com
}

func (i Invoker) executeComand(com Command) {
	if i.c != nil {
		i.c.execute()
	}
}

func (i Invoker) undoComand(com Command) {
	if i.c != nil {
		i.c.undo()
	}
}

func main() {
	receiver := Receiver{}
	command := ConcreteCommand(receiver){}
}
