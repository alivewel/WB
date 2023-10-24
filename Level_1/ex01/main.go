package main

import "fmt"

// Определяем структуру Human
type Human struct {
	Name  string
	Age   int
}

// Метод для структуры Human
func (h *Human) SayHello() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

// Определяем структуру Action, встраивая Human
type Action struct {
	Human
	ActionDescription string
}

func main() {
	// Создаем экземпляр структуры Action
	person := Action{
		Human: Human{
			Name:  "Иван",
			Age:   30,
		},
		ActionDescription: "Выполнение действия",
	}

	// Используем методы и поля как будто они принадлежат структуре Action
	fmt.Println(person.Name)              // Иван
	fmt.Println(person.ActionDescription) // Выполняет какое-то действие
	person.SayHello()                     // Вызываем метод SayHello() структуры Human
}
