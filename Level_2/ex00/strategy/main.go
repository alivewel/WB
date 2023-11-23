package main

func main() {
	cat := newCat()
	dog := newDog()

	context := &Context{}
	context.SetAnimal(cat)
	context.MakeAnimalSound()

	context.SetAnimal(dog)
	context.MakeAnimalSound()
}

// Команда для запуска:
// go run .
