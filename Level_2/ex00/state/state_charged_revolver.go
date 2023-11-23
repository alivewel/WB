package main

import "fmt"

type StateChargedRevolver struct {
	revolver *Revolver
}

func (s *StateChargedRevolver) charge() error {
	if s.revolver.patronCount != 6 {
		s.revolver.patronCount = 6
		s.revolver.setState(s.revolver.charged)
		return nil
	} else {
		return fmt.Errorf("The revolver is already loaded!")
	}
}

func (s *StateChargedRevolver) shoot() error {
	if s.revolver.patronCount > 0 {
		fmt.Println("Shoot!")
		s.revolver.incrementPatron()
		return nil
	} else {
		s.revolver.setState(s.revolver.discharged)
		return fmt.Errorf("Patrons are out!")
	}
}

func (s *StateChargedRevolver) switchFuse() error {
	s.revolver.setState(s.revolver.onFuse)
	fmt.Println("The revolver is on safety!")
	return nil
}
