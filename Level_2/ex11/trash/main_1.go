package main

import (
	"fmt"
	"time"
)

func main() {
	// Пример: преобразование time.Time в строку
	currentTime := time.Now()
	// timeString := currentTime.Format("2006-01-02 15:04:05")
	timeString := currentTime.Format("2006-01-02")

	fmt.Println(timeString)
}
