package main

import (
	"fmt"
	"sync"
)

// с использованием sync.Map
func main() {
	var m sync.Map
	var wg sync.WaitGroup
	// запись
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key, value int) {
			m.Store(key, value)
			wg.Done()
		}(i, i*i)
	}

	wg.Wait()

	// чтение
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(key int) {
			value, ok := m.Load(key)
			if ok {
				fmt.Printf("Ключ: %d, Значение: %v\n", key, value)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
