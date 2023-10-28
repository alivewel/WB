package main

import (
	"fmt"
	"time"
)

func worker(stop <-chan bool) {
	for {
		select {
		default:
			// Выполнение работы в горутине
			fmt.Println("Работаю...")
			time.Sleep(1 * time.Second)
		case <-stop:
			// Получен сигнал об остановке
			fmt.Println("Остановлен")
			return
		}
	}
}

func main() {
	// Создание канала для передачи сигналов об остановке горутины
	stop := make(chan bool)

	// Запуск горутины
	go worker(stop)

	// Ждем 5 секунд, затем отправляем сигнал об остановке
	time.Sleep(5 * time.Second)
	stop <- true

	// Ждем завершения горутины
	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена")
}
