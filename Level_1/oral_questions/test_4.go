package main

import "fmt"

// Интерфейс, который определяет метод, принимающий int как аргумент
type ArgumentInterface interface {
	Process(int)
}

// Структура, которая удовлетворяет интерфейсу ArgumentInterface
type IntProcessor struct{}

func (p IntProcessor) Process(value int) {
	fmt.Printf("Обработка int: %d\n", value)
}

// Функция, принимающая аргумент через интерфейс
func AcceptArgument(arg ArgumentInterface, value int) {
	arg.Process(value)
}

func main() {
	processor := IntProcessor{}   // Создаем экземпляр структуры
	AcceptArgument(processor, 42) // Передаем аргумент через интерфейс
}
