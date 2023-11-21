package main

type Waitress struct {
	order Order
}

func (w *Waitress) SetOrder(order Order) {
	w.order = order
}

func (w *Waitress) ExecuteOrder() {
	w.order.Execute()
}
