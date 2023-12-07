package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		select {
		case <-channels[0]:
		case <-channels[1]:
		case <-channels[2]:
			// Добавьте необходимое количество case для других каналов
		}
	}()
	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v", time.Since(start))
}
