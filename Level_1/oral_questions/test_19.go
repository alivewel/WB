package main

import "fmt"

// Интерфейс для создания кнопок
type Button interface {
	Paint()
}

// Интерфейс для создания фабрики кнопок
type ButtonFactory interface {
	CreateButton() Button
}

// Конкретная реализация кнопки для Windows
type WindowsButton struct{}

func (w *WindowsButton) Paint() {
	fmt.Println("Rendering a Windows button")
}

// Конкретная реализация фабрики кнопок для Windows
type WindowsButtonFactory struct{}

func (wbf WindowsButtonFactory) CreateButton() Button {
	return &WindowsButton{}
}

// Конкретная реализация кнопки для macOS
type MacOSButton struct{}

func (m MacOSButton) Paint() {
	fmt.Println("Rendering a macOS button")
}

// Конкретная реализация фабрики кнопок для macOS
type MacOSButtonFactory struct{}

func (mbf MacOSButtonFactory) CreateButton() Button {
	return &MacOSButton{}
}

func main() {
	// Использование абстрактной фабрики для создания кнопок
	windowsFactory := WindowsButtonFactory{}
	macOSFactory := MacOSButtonFactory{}

	windowsButton := windowsFactory.CreateButton()
	macOSButton := macOSFactory.CreateButton()

	// Отрисовка кнопок
	windowsButton.Paint()
	macOSButton.Paint()
}
