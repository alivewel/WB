package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; ; i++ {
			i++
		}
	}()
	<-time.After(time.Second * 1)
	fmt.Println("Hello world!")
}
