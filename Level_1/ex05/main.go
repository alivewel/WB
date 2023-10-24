package main

import (
	"fmt"
	"time"
)

func sender(c chan int) {
	for i := 1; i <= 10; i++ {
		c <- i
		time.Sleep(1 * time.Second) // Отправляем значение каждую секунду
	}
	close(c)
}

func main() {
	startTime := time.Now() // Засекаем время начала работы программы

	c := make(chan int)

	go sender(c)

	N := 5 // Установите значение N в секундах

	timer := time.NewTimer(time.Duration(N) * time.Second)
	defer timer.Stop()

	for {
		select {
		case val, ok := <-c:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Принято:", val)
		case <-timer.C:
			fmt.Printf("Прошло %d секунд, программа завершается\n", N)
			endTime := time.Now() // Засекаем время завершения программы
			elapsedTime := endTime.Sub(startTime)
			fmt.Printf("Время работы программы: %s\n", elapsedTime)
			return
		}
	}
}
