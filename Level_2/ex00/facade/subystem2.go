package main

import "fmt"

type SubSystem2 struct{}

func (s SubSystem2) Suboperation() {
	fmt.Println("SubSystem2 operation")
}
