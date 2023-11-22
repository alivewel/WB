package main

import "fmt"

type Location interface {
	accept(Visitor)
}

type Home struct {
	action string
}

func (h *Home) accept(v Visitor) {
	v.visitHome(h)
}

func newHome() *Home {
	return &Home{action: "Sleep"}
}

type Work struct {
	action string
}

func (w *Work) accept(v Visitor) {
	v.visitWork(w)
}

func newWork() *Work {
	return &Work{action: "Do tasks"}
}

type School struct {
	action string
}

func (s *School) accept(v Visitor) {
	v.visitSchool(s)
}

func newSchool() *School {
	return &School{action: "Study"}
}

type Visitor interface {
	visitHome(*Home)
	visitWork(*Work)
	visitSchool(*School)
}

type VisitHome struct{}

func (v *VisitHome) visitHome(h *Home) {
	fmt.Println(h.action)
}

type VisitWork struct{}

func (v *VisitHome) visitWork(w *Work) {
	fmt.Println(w.action)
}

type VisitSchool struct{}

func (v *VisitHome) visitSchool(s *School) {
	fmt.Println(s.action)
}

func main() {

}

// Команда для запуска:
// go run .
