package main

import "fmt"

func main() {
	can := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			can <- i
		}
		close(can)
	}()

	for val := range can {
		fmt.Println(val)
	}
	fmt.Println(<-can)
}
