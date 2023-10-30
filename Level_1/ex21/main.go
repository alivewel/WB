package main

import "fmt"

// OldPrinter - структура старой системы для печати
type OldPrinter struct {
	msg string
}

func (p *OldPrinter) printMessage() {
	fmt.Println(p.msg)
}

// NewPrinter - структура новой системы для печати
type NewPrinter struct {
	message string
}

func (p *NewPrinter) print() {
	fmt.Println(p.message)
}

// PrinterAdapter - адаптер, позволяющий использовать OldPrinter с интерфейсом NewPrinter
type PrinterAdapter struct {
	oldPrinter *OldPrinter
	message    string
}

func (pa *PrinterAdapter) print() {
	pa.oldPrinter.msg = pa.message
	pa.oldPrinter.printMessage()
}

func main() {
	newPrinter := &NewPrinter{message: "Это новая система печати."}
	oldPrinter := &OldPrinter{msg: "Это старая система печати."}

	// Используем новую систему печати
	newPrinter.print()

	// Используем адаптер, чтобы использовать старую систему печати через интерфейс новой системы
	adapter := &PrinterAdapter{oldPrinter: oldPrinter, message: "Адаптер использует старую систему."}
	adapter.print()
}
