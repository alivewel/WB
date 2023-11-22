package main

import "fmt"

type Location interface {
	accept(Visitor)
}

type Home struct {
	action string
	want   string
}

func (h *Home) accept(v Visitor) {
	v.visitHome(h)
}

func newHome() *Home {
	return &Home{
		action: "Go to sleep",
		want:   "Want to go home"}
}

type Work struct {
	action string
	want   string
}

func (w *Work) accept(v Visitor) {
	v.visitWork(w)
}

func newWork() *Work {
	return &Work{
		action: "Do tasks",
		want:   "Want to receive money"}
}

type School struct {
	action string
	want   string
}

func (s *School) accept(v Visitor) {
	v.visitSchool(s)
}

func newSchool() *School {
	return &School{
		action: "Starting to study",
		want:   "Want to know more"}
}

type Visitor interface {
	visitHome(*Home)
	visitWork(*Work)
	visitSchool(*School)
}

type VisitorAction struct{}

func (v *VisitorAction) visitHome(h *Home) {
	fmt.Println(h.action)
}

func (v *VisitorAction) visitWork(w *Work) {
	fmt.Println(w.action)
}

func (v *VisitorAction) visitSchool(s *School) {
	fmt.Println(s.action)
}

type VisitorWant struct{}

func (v *VisitorWant) visitHome(h *Home) {
	fmt.Println(h.want)
}

func (v *VisitorWant) visitWork(w *Work) {
	fmt.Println(w.want)
}

func (v *VisitorWant) visitSchool(s *School) {
	fmt.Println(s.want)
}

func main() {
	home := newHome()
	work := newWork()
	school := newSchool()

	visitorAction := VisitorAction{}

	visitorAction.visitHome(home)
	visitorAction.visitWork(work)
	visitorAction.visitSchool(school)

	visitorWant := VisitorWant{}

	visitorWant.visitHome(home)
	visitorWant.visitWork(work)
	visitorWant.visitSchool(school)
}

// Команда для запуска:
// go run .
