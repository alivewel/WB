package main

func main() {
	cook := &Cook{}

	makeBurger := &OrderMakeBurger{cook: cook}
	makeFries := &OrderMakeFries{cook: cook}

	waitress := &Waitress{}

	waitress.SetOrder(makeBurger)
	waitress.ExecuteOrder()

	waitress.SetOrder(makeFries)
	waitress.ExecuteOrder()
}

// Команда для запуска:
// go run .
