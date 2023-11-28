package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type Flags struct {
	flagA int
	flagB int
	flagC int
	flagc bool
	flagI bool
	flagV bool
	flagF bool
	flagN bool
}

var flags Flags

func init() {
	flag.IntVar(&flags.flagA, "A", 0, "Print N lines after each match")
	flag.IntVar(&flags.flagB, "B", 0, "Print N lines before each match")
	flag.IntVar(&flags.flagC, "C", 0, "Print N lines of output context")
	flag.BoolVar(&flags.flagc, "c", false, "Print only the number of matches")
	flag.BoolVar(&flags.flagI, "i", false, "Case-insensitive match")
	flag.BoolVar(&flags.flagV, "v", false, "Select non-matching lines")
	flag.BoolVar(&flags.flagF, "F", false, "Fixed string (disable regular expressions)")
	flag.BoolVar(&flags.flagN, "n", false, "Show line numbers")
	flag.Parse()
}

func main() {
	// Парсинг флагов
	fmt.Println("Flag A:", flags.flagA)
	fmt.Println("Flag B:", flags.flagB)
	fmt.Println("Flag C:", flags.flagC)
	fmt.Println("Flag c:", flags.flagc)
	fmt.Println("Flag I:", flags.flagI)
	fmt.Println("Flag V:", flags.flagV)
	fmt.Println("Flag F:", flags.flagF)
	fmt.Println("Flag N:", flags.flagN)

	// Получение остальных аргументов (паттерн и файлы)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: grep-go [OPTIONS] PATTERN FILE [FILE...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	fmt.Println("files", files)
	fmt.Println("pattern", pattern)

	// Применение фильтрации к каждому файлу
	// for _, file := range files {
	// 	err := grepFile(pattern, file, *afterFlag, *beforeFlag, *contextFlag)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 	}
	// }
}

func grepFile(pattern, filename string, after, before, context int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Проверка на совпадение паттерна
		if containsPattern(line, pattern) {
			fmt.Println(line)
		}
	}

	return nil
}

func containsPattern(line, pattern string) bool {

	return true // Ваша реализация проверки наличия паттерна в строке
}
