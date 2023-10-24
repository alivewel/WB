package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, wg *sync.WaitGroup, dataChan <-chan int) {
	defer wg.Done()

	for data := range dataChan {
		fmt.Printf("Worker %d received: %d\n", id, data)
		time.Sleep(time.Second) // Добавляем задержку в 1 секунду
	}
}

func main() {
	numWorkers := flag.Int("w", 5, "An integer value")
	flag.Parse()

	if *numWorkers < 1 {
		log.Fatal("The number of workers must be greater than or equal to 1")
	}

	dataChan := make(chan int)
	var wg sync.WaitGroup

	// Запускаем воркеров
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg, dataChan)
	}

	// Обработка сигнала Ctrl+C для завершения работы
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		close(dataChan) // Закрываем канал при получении сигнала
	}()

	// Записываем произвольные данные в канал (главный поток)
	for i := 1; i <= 10; i++ {
		dataChan <- i
	}

	close(dataChan) // Закрываем канал после завершения записи данных

	wg.Wait() // Ожидаем завершения всех воркеров

	fmt.Println("Main goroutine stopped")
}

// go run main.go -w 1
