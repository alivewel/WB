package main

import (
	"fmt"
)

func createNum(c1 chan int, arrayInt []int) {
	for i := 0; i < len(arrayInt); i++ {
		c1 <- arrayInt[i]
	}
	close(c1) // close channel
}

func squares(c1, c2 chan int, arrayInt []int) {
	for i := 0; i < len(arrayInt); i++ {
		res := <-c1
		c2 <- res * res
	}
	close(c2) // close channel
}

func main() {
	fmt.Println("main() started")
	c1 := make(chan int)
	c2 := make(chan int)
	arrayInt := []int{5, 1, 3, 9, 4, 6}
	go createNum(c1, arrayInt) // start goroutine

	go squares(c1, c2, arrayInt)

	for {
		val, ok := <-c2
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break // exit break loop
		} else {
			fmt.Println(val, ok)
		}
	}


	fmt.Println("main() stopped")
}
