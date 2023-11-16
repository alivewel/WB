package main

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func main() {
	currentTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Текущее время:", currentTime)
}
