package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d started\n", id)

			ctx, cancel := context.WithCancel(context.Background())

			if id == 1 {
				// Закрываем контекст для завершения горутины с идентификатором 1
				fmt.Printf("Goroutine %d \n", id)
				cancel()
			}

			select {
			case <-ctx.Done():
				fmt.Printf("Goroutine %d canceled\n", id)
				return
			default:
				// Выполняем работу горутины
				fmt.Printf("Goroutine %d finished\n", id)
			}
		}(i)
	}

	wg.Wait()
}
