package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d started\n", id)

			if id == 1 {
				// Завершаем выполнение текущей горутины
				runtime.Goexit()
			}

			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	wg.Wait()
}

// go run test_9.go
// Goroutine 2 started
// Goroutine 2 finished
// Goroutine 0 started
// Goroutine 0 finished
// Goroutine 1 started
