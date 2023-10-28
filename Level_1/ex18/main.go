package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    sync.Mutex
}

func (c *counter) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	count := new(counter)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go count.Increment(&wg)
	}
	wg.Wait()
	fmt.Println(count.count)
}

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
