package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwm sync.RWMutex
	mapInt := make(map[int]int, 0)
	var wg sync.WaitGroup
	// запись
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(m *sync.RWMutex, key, value int) {
			m.Lock() // write-lock
			mapInt[key] = value
			m.Unlock() // write-unlock
			wg.Done()
		}(&rwm, i, i*i)
	}

	wg.Wait()

	// чтение
	for key, value := range mapInt {
		wg.Add(1)
		go func(m *sync.RWMutex, key, value int) {
			m.Lock()
			fmt.Printf("Ключ: %d, Значение: %d\n", key, value)
			m.Unlock()
			wg.Done()
		}(&rwm, key, value)
	}

	wg.Wait()
}
