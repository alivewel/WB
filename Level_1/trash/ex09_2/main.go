package main

import (
	"fmt"
	// "time"
)

func createNum(c1 chan int) {
	for i := 1; i <= 9; i++ {
		c1 <- i
	}
	close(c1) // close channel
}

func squares(c1, c2 chan int) {
	for i := 1; i <= 9; i++ {
		res := <-c1
		c2 <- res * res
	}
	close(c2) // close channel
}

func main() {
	fmt.Println("main() started")
	c1 := make(chan int)
	c2 := make(chan int)
	// arrayInt := []int{5, 1, 3, 9, 4, 6}
	go createNum(c1) // start goroutine

	go squares(c1, c2)

	for {
		val, ok := <-c2
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break // exit break loop
		} else {
			fmt.Println(val, ok)
		}
	}

	// time.Sleep(1 * time.Second)

	fmt.Println("main() stopped")
}
