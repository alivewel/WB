package main

import "fmt"

type StateOnFuseRevolver struct {
	revolver *Revolver
}

func (s *StateOnFuseRevolver) charge() error {
	if s.revolver.patronCount != 6 {
		s.revolver.patronCount = 6
		s.revolver.setState(s.revolver.charged)
		return nil
	} else {
		return fmt.Errorf("The revolver is already loaded!")
	}
}

func (s *StateOnFuseRevolver) shoot() error {
	return fmt.Errorf("The revolver is on safety!")
}

func (s *StateOnFuseRevolver) switchFuse() error {
	if s.revolver.patronCount > 0 {
		s.revolver.setState(s.revolver.charged)
		fmt.Println("The revolver is off safety!")
		return nil
	} else {
		s.revolver.setState(s.revolver.discharged)
		fmt.Println("The revolver is off safety!")
		return nil
	}
}
