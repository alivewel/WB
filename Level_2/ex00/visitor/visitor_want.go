package main

import "fmt"

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
