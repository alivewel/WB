package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

type Action struct {
	Human
	ActionDescription string
}

func main() {
	person := Action{
		Human: Human{
			Name: "Иван",
			Age:  30,
		},
		ActionDescription: "Выполнение действия",
	}

	fmt.Println(person.Name)
	fmt.Println(person.ActionDescription)
	person.SayHello()
}
