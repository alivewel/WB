package main

import (
	"fmt"
)

func main() {
	revolver := newRevolver(1)

	err := revolver.shoot()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	err = revolver.charge()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	err = revolver.switchFuse()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	err = revolver.shoot()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

// Команда для запуска:
// go run .
