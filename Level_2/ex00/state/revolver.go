package main

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
