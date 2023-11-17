package main

import "fmt"

type SubSystem1 struct{}

func (s SubSystem1) Suboperation() {
	fmt.Println("SubSystem1 operation")
}

type SubSystem2 struct{}

func (s SubSystem2) Suboperation() {
	fmt.Println("SubSystem2 operation")
}

type Facade struct {
	subSystem1 SubSystem1
	subSystem2 SubSystem2
}

func (f Facade) OperationWrapper() {
	f.subSystem1.Suboperation()
	f.subSystem2.Suboperation()
}

func main() {
	facade := Facade{}
	facade.OperationWrapper()
}
