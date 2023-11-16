package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	// Вызов системного вызова fork
	pid, err := syscall.ForkExec(os.Args[0], os.Args, nil)
	if err != nil {
		fmt.Printf("Ошибка при вызове fork: %v\n", err)
		return
	}

	if pid == 0 {
		// Этот код будет выполняться в дочернем процессе
		fmt.Fprintln(os.Stdout, "Дочерний процесс")
	} else {
		// Этот код будет выполняться в родительском процессе
		fmt.Fprintln(os.Stdout, "Родительский процесс")
		// Добавим небольшую задержку, чтобы увидеть вывод от дочернего процесса
		time.Sleep(time.Second)
	}
}
