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

func (r *Revolver) switchFuse() error {
	return r.currentState.switchFuse()
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
	switchFuse() error
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
		s.revolver.setState(s.revolver.discharged)
		return fmt.Errorf("Patrons are out!")
	}
}

func (s *StateChargedRevolver) switchFuse() error {
	s.revolver.setState(s.revolver.onFuse)
	fmt.Println("The revolver is on safety!")
	return nil
}

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

	err = revolver.switchFuse()
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
