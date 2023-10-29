package main

import (
	"fmt"
)

// с использованием буфиризованных каналов

func createNum(c1 chan int, arrayInt []int) {
	for i := 0; i < len(arrayInt); i++ {
		c1 <- arrayInt[i]
	}
	close(c1)
}

func squares(c1, c2 chan int) {
	for i := range c1 {
		c2 <- i * 2
	}
	close(c2)
}

func main() {
	arrayInt := []int{5, 1, 3, 9, 4, 6}
	c1 := make(chan int, len(arrayInt))
	c2 := make(chan int, len(arrayInt))
	go createNum(c1, arrayInt)

	go squares(c1, c2)

	for {
		val, ok := <-c2
		if ok == false {
			break
		} else {
			fmt.Println(val)
		}
	}
}
