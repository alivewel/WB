package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// os.Args содержит срез строк, представляющих аргументы командной строки
	args := os.Args

	ExecuteCommand(args)

	// path := "/Users/alivewel/WB/Level_2"

	// // Используем функцию filepath.Dir() для отсечения последнего отрезка
	// parentDir := filepath.Dir(path)

	// fmt.Println(parentDir)
	// Последующие элементы - аргументы командной строки
	// fmt.Println("Аргументы командной строки:", args[1:])
}

func ExecuteCommand(args []string) {
	if len(args) > 1 {
		nameUtility := args[1]
		switch nameUtility {
		case "cd":
			cd(args)
		case "pwd":
			pwd(args)
		case "echo":
			echo()
		case "kill":
			kill()
		case "ps":
			ps()
		}
	}
}

// директория меняется внутри программы, но в терминале остается прежней
func cd(args []string) {
	if len(args) == 3 {
		newDir := args[2]
		if newDir == ".." { // команда cd ..
			currentDir, _ := os.Getwd()
			newDir = filepath.Dir(currentDir)
		}
		err := os.Chdir(newDir)
		if err != nil {
			fmt.Println("Ошибка при переходе в директорию:", err)
			os.Exit(1)
		}
		// cwd, _ := os.Getwd()
		// fmt.Println("cwd:", cwd)
	}
}

func pwd(args []string) {
	if len(args) == 2 {
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Ошибка при переходе в директорию:", err)
			os.Exit(1)
		}
		fmt.Println(currentDir)
	}
}

func echo() {

}

func kill() {

}

func ps() {

}
