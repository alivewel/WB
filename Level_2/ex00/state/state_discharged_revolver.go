package main

import "fmt"

type StateDischargedRevolver struct {
	revolver *Revolver
}

func (s *StateDischargedRevolver) charge() error {
	s.revolver.patronCount = 6
	s.revolver.setState(s.revolver.charged)
	return nil
}

func (s *StateDischargedRevolver) shoot() error {
	return fmt.Errorf("Patrons are out!!!")
}

func (s *StateDischargedRevolver) switchFuse() error {
	s.revolver.setState(s.revolver.onFuse)
	fmt.Println("The revolver is on safety!")
	return nil
}
