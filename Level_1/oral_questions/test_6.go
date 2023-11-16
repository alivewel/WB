package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(os.Args[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Ошибка при запуске дочернего процесса: %v\n", err)
	}
}
