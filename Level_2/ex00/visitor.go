package main

import "fmt"

// Интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element ConcreteElementA)
	VisitConcreteElementB(element ConcreteElementB)
}

// Конкретный посетитель A
type ConcreteVisitorA struct{}

func (v ConcreteVisitorA) VisitConcreteElementA(element ConcreteElementA) {
	fmt.Println("ConcreteVisitorA посещает", element.GetName())
}

func (v ConcreteVisitorA) VisitConcreteElementB(element ConcreteElementB) {
	fmt.Println("ConcreteVisitorA посещает", element.GetName())
}

// Конкретный посетитель B
type ConcreteVisitorB struct{}

func (v ConcreteVisitorB) VisitConcreteElementA(element ConcreteElementA) {
	fmt.Println("ConcreteVisitorB посещает", element.GetName())
}

func (v ConcreteVisitorB) VisitConcreteElementB(element ConcreteElementB) {
	fmt.Println("ConcreteVisitorB посещает", element.GetName())
}

// Интерфейс элемента
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type ConcreteElementA struct {
	name string
}

func (e ConcreteElementA) GetName() string {
	return e.name
}

func (e ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// Конкретный элемент B
type ConcreteElementB struct {
	name string
}

func (e ConcreteElementB) GetName() string {
	return e.name
}

func (e ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Структура объектов
type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) AddElement(element Element) {
	os.elements = append(os.elements, element)
}

func (os ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

func main() {
	objectStructure := ObjectStructure{}
	objectStructure.AddElement(ConcreteElementA{"ElementA"})
	objectStructure.AddElement(ConcreteElementB{"ElementB"})

	visitorA := ConcreteVisitorA{}
	visitorB := ConcreteVisitorB{}

	objectStructure.Accept(visitorA)
	objectStructure.Accept(visitorB)
}
