package main

import (
	"fmt"
	"time"
)

func worker(stop <-chan bool) {
	for {
		select {
		default:
			fmt.Println("Работаю...")
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("Остановлен")
			return
		}
	}
}

func main() {
	stop := make(chan bool)

	go worker(stop)

	time.Sleep(5 * time.Second)
	stop <- true

	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена")
}
