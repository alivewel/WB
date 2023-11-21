package main

import "fmt"

type SubSystem1 struct{}

func (s SubSystem1) Suboperation() {
	fmt.Println("SubSystem1 operation")
}
