package main

import "fmt"

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
