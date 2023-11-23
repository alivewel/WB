package main

import (
	"fmt"
	"log"
)

type Revolver struct {
	charged    State
	discharged State
	onFuse     State

	currentState State

	patronCount int
}

func newRevolver(patronCount int) *Revolver {
	r := &Revolver{patronCount: patronCount}

	chargedRevolver := &StateChargedRevolver{revolver: r}
	dischargedRevolver := &StateDischargedRevolver{revolver: r}
	onFuseRevolver := &StateOnFuseRevolver{revolver: r}

	r.setState(chargedRevolver)
	r.charged = chargedRevolver
	r.discharged = dischargedRevolver
	r.onFuse = onFuseRevolver
	return r
}

func (r *Revolver) setState(s State) {
	r.currentState = s
}

func (r *Revolver) charge() error {
	return r.currentState.charge()
}

func (r *Revolver) shoot() error {
	return r.currentState.shoot()
}

func (r *Revolver) putOnFuse() error {
	return r.currentState.putOnFuse()
}

func (r *Revolver) incrementPatron() {
	if r.patronCount > 0 {
		r.patronCount = r.patronCount - 1
	}
}

func (r *Revolver) getPatronCount() int {
	return r.patronCount
}

type State interface {
	charge() error
	shoot() error
	putOnFuse() error
}

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
		s.revolver.setState(s.revolver.charged)
		return fmt.Errorf("Patrons are out!")
	}
}

func (s *StateChargedRevolver) putOnFuse() error {
	if s.revolver.currentState != s.revolver.discharged {
		s.revolver.setState(s.revolver.discharged)
		fmt.Println("The revolver is on safety!")
		return nil
	} else {
		return fmt.Errorf("The revolver is already loaded!")
	}
}

type StateDischargedRevolver struct {
	revolver *Revolver
}

func (s *StateDischargedRevolver) charge() error {
	if s.revolver.patronCount != 6 {
		s.revolver.patronCount = 6
		s.revolver.setState(s.revolver.charged)
		return nil
	} else {
		return fmt.Errorf("The revolver is already loaded!")
	}
}

func (s *StateDischargedRevolver) shoot() error {
	return fmt.Errorf("Patrons are out!")
}

func (s *StateDischargedRevolver) putOnFuse() error {
	if s.revolver.currentState != s.revolver.discharged {
		s.revolver.setState(s.revolver.discharged)
		fmt.Println("The revolver is on safety!")
		return nil
	} else {
		return fmt.Errorf("The revolver is already loaded!")
	}
}

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

func (s *StateOnFuseRevolver) putOnFuse() error {
	return fmt.Errorf("The revolver is already on safety!")
}

func main() {
	revolver := newRevolver(1)

	err := revolver.shoot()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(revolver.getPatronCount())

	err = revolver.charge()
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(revolver.getPatronCount())

	err = revolver.putOnFuse()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = revolver.shoot()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

// Команда для запуска:
// go run .
