package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
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
	// fmt.Println("Flag A:", flags.flagA)
	// fmt.Println("Flag B:", flags.flagB)
	// fmt.Println("Flag C:", flags.flagC)
	// fmt.Println("Flag c:", flags.flagc)
	// fmt.Println("Flag I:", flags.flagI)
	// fmt.Println("Flag V:", flags.flagV)
	// fmt.Println("Flag F:", flags.flagF)
	// fmt.Println("Flag N:", flags.flagN)

	// Получение остальных аргументов (паттерн и файлы)
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: grep-go [OPTIONS] PATTERN FILE [FILE...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	// fmt.Println("files", files)
	// fmt.Println("pattern", pattern)

	var allScanStr [][]string
	var allIndexStr [][]int
	// Применение фильтрации к каждому файлу
	for _, file := range files {
		scanStr, indexStr, err := grepFile(pattern, file)
		if err != nil {
			fmt.Println("Error:", err)
		}
		allScanStr = append(allScanStr, scanStr)
		allIndexStr = append(allIndexStr, indexStr)
	}

	for i := 0; i < len(allScanStr); i++ {
		for j := 0; j < len(allScanStr[i]); j++ {
			fmt.Println(allScanStr[i][j])
		}
		fmt.Println()
	}
}

func grepFile(pattern, filename string) ([]string, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var scanStr []string
	var indexStr []int
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		scanStr = append(scanStr, line)
		count++
		// Проверка на совпадение паттерна
		if containsPattern(line, pattern) {
			fmt.Println(line)
			// indexStr = append(indexStr, count)
			addRangeNum(count, &indexStr)
		}
	}

	return scanStr, indexStr, nil
}

func containsPattern(line, pattern string) bool {
	if flags.flagI {
		line = strings.ToLower(line)
		line = strings.ToLower(pattern)
	}
	if flags.flagF {
		return strings.Contains(line, pattern)
	} else {
		re := regexp.MustCompile(pattern)
		return re.MatchString(line)
	}
}

func generateRange(start, offset int, result *[]int) {
	if offset > 0 {
		// В случае положительного offset
		for i := start + 1; i <= start+offset; i++ {
			*result = append(*result, i)
		}
	} else if offset < 0 {
		// В случае отрицательного offset
		for i := start - 1; i >= start+offset; i-- {
			if i < 0 {
				break
			}
			*result = append(*result, i)
		}
	} else {
		// В случае offset равного нулю
		*result = append(*result, start)
	}
}

func addRangeNum(start int, result *[]int) {
	if flags.flagA > 0 || flags.flagC > 0 {
		offset := max(flags.flagA, flags.flagC)
		generateRange(start, offset, result)
	}
	if flags.flagB > 0 || flags.flagC > 0 {
		offset := max(flags.flagB, flags.flagC)
		generateRange(start, -offset, result)
	}
	*result = append(*result, start)
}

func removeDuplicates(input []int) []int {
	uniqueMap := make(map[int]bool)
	var result []int
	for _, num := range input {
		if !uniqueMap[num] {
			// Если элемент не встречался ранее, добавляем его в результат и карту
			result = append(result, num)
			uniqueMap[num] = true
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
