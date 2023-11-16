package main

import (
	"fmt"
	"log"
	"os"
	"time"

	fork "github.com/kraken-hpc/go-fork"
)

func init() {
	fork.RegisterFunc("child", child)
	fork.Init()
}

// func child(n int) {
// 	fmt.Printf("child(%d) pid: %d\n", n, os.Getpid())
// }

func child(n int) {
	// Открываем файл для записи вывода дочернего процесса
	outputFile, err := os.Create(fmt.Sprintf("child_output_%d.txt", n))
	if err != nil {
		log.Fatalf("Ошибка при открытии файла для вывода: %v", err)
	}
	defer outputFile.Close()

	// Перенаправляем стандартный вывод в файл
	log.SetOutput(outputFile)

	fmt.Printf("child(%d) pid: %d\n", n, os.Getpid())
}

func main() {
	fmt.Printf("main() pid: %d\n", os.Getpid())
	if err := fork.Fork("child", 1); err != nil {
		log.Fatalf("failed to fork: %v", err)
	}
	time.Sleep(1 * time.Second)
}

// https://pkg.go.dev/github.com/kraken-hpc/go-fork@v0.1.1/example
