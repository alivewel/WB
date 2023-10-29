package main

import (
	"fmt"
	"time"
)

func createNum(c1 chan int) {
	// for i := 0; i <= 9; i++ {
	// 	c1 <- i
	// }
	c1 <- 2
	fmt.Println("createNum")
	close(c1) // close channel
}

func squares(c1, c2 chan int) {
	// for i := 0; i <= 9; i++ {
	// 	res := <- c1
	// 	c2 <- res * res
	// }
	fmt.Println("squares")
	res := <-c1
	c2 <- res * res

	close(c2) // close channel
}

func main() {
	fmt.Println("main() started")
	c1 := make(chan int)
	c2 := make(chan int)

	go createNum(c1) // start goroutine
	go squares(c1, c2)

	res := <-c2

	fmt.Println("res", res)
	// periodic block/unblock of main goroutine until chanel closes
	// for {
	// 	val, ok := <-c
	// 	if ok == false {
	// 		fmt.Println(val, ok, "<-- loop broke!")
	// 		break // exit break loop
	// 	} else {
	// 		fmt.Println(val, ok)
	// 	}
	// }

	time.Sleep(1 * time.Second)

	fmt.Println("main() stopped")
}
