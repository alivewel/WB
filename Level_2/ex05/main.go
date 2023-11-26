package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Парсинг флагов
	afterFlag := flag.Int("A", 0, "Print N lines after each match")
	beforeFlag := flag.Int("B", 0, "Print N lines before each match")
	contextFlag := flag.Int("C", 0, "Print N lines of output context")
	flag.Parse()

	// Получение остальных аргументов (паттерн и файлы)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: grep-go [OPTIONS] PATTERN FILE [FILE...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	fmt.Println("afterFlag", *afterFlag)
	fmt.Println("beforeFlag", *beforeFlag)
	fmt.Println("contextFlag", *contextFlag)

	// fmt.Println("files", files)
	// fmt.Println("pattern", pattern)

	// Применение фильтрации к каждому файлу
	for _, file := range files {
		err := grepFile(pattern, file, *afterFlag, *beforeFlag, *contextFlag)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func grepFile(pattern, filename string, after, before, context int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matchingLines := make([]string, 0)
	contextLines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		// Проверка на совпадение паттерна
		if containsPattern(line, pattern) {
			// Выводим строки до совпадения (если есть)
			for _, beforeLine := range contextLines {
				fmt.Println(beforeLine)
			}
			contextLines = contextLines[:0]

			// Выводим строки совпадения
			fmt.Println(line)

			// Выводим строки после совпадения (если есть)
			for i := 1; i <= after; i++ {
				if scanner.Scan() {
					fmt.Println(scanner.Text())
				}
			}
			matchingLines = matchingLines[:0]
		} else {
			// Если не совпало, добавляем строку в контекст
			contextLines = append(contextLines, line)

			// Поддержка флага -C (контекст)
			if len(contextLines) > context {
				contextLines = contextLines[1:]
			}
		}

		// Поддержка флага -B (перед совпадением)
		if len(matchingLines) > before {
			matchingLines = matchingLines[1:]
		}
		matchingLines = append(matchingLines, line)
	}

	return nil
}

func containsPattern(line, pattern string) bool {
	return true // Ваша реализация проверки наличия паттерна в строке
}
