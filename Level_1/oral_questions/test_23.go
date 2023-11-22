package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	// Горутина-писатель
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	// Горутина-читатель
	wg.Add(1)
	go func() {
		defer wg.Done()
		for value := range ch {
			fmt.Println(value)
		}
	}()

	// Ждем завершения всех горутин
	wg.Wait()
}
