package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
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

// тесты с функцией init не работают, пришлось создать свою функцию
func initFlags() {
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
	initFlags()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: grep-go [OPTIONS] PATTERN FILE [FILE...]")
		os.Exit(1)
	}

	pattern := args[0]
	files := args[1:]

	var allScanStr [][]string // все отсканированные строки в каждом файле
	var allIndexStr [][]int   // индексы совпадающих строк в каждом файле
	var countContain []int    // количество совпадений в каждом файле

	for _, file := range files {
		var count int
		scanStr, indexStr, err := grepFile(pattern, file, &count)
		if err != nil {
			fmt.Println("Error:", err)
		}
		allScanStr = append(allScanStr, scanStr)
		allIndexStr = append(allIndexStr, indexStr)
		countContain = append(countContain, count)
	}

	prepareIndex(allIndexStr)

	// инвертируем все индексы в массиве allIndexStr
	if flags.flagV {
		allIndexStr = invertIndex(allIndexStr, allScanStr)
	}

	printResult(allScanStr, allIndexStr, countContain, files)
}

func grepFile(pattern, filename string, countContain *int) ([]string, []int, error) {
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

		// Проверка на совпадение паттерна
		if containsPattern(line, pattern) {
			addRangeNum(count, &indexStr)
			*countContain++
		}
		count++
	}

	return scanStr, indexStr, nil
}

func containsPattern(line, pattern string) bool {
	if flags.flagI {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
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
		for i := start; i >= start+offset; i-- {
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

func prepareIndex(input [][]int) {
	for i := 0; i < len(input); i++ {
		input[i] = removeDuplicates(input[i])
		sort.Ints(input[i])
	}
}

func invertIndex(allIndexStr [][]int, allScanStr [][]string) [][]int {
	result := make([][]int, len(allScanStr))
	for i := 0; i < len(allScanStr); i++ {
		for j := 0; j < len(allScanStr[i]); j++ {
			if !containsElemArr(allIndexStr[i], j) {
				result[i] = append(result[i], j)
			}
		}
	}
	return result
}

func containsElemArr(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printResult(allScanStr [][]string, allIndexStr [][]int, countContain []int, files []string) {
	if flags.flagc {
		for i := 0; i < len(allIndexStr); i++ {
			if len(allIndexStr) > 1 {
				fmt.Printf("%s: ", files[i])
			}
			fmt.Println(countContain[i])
		}
	} else {
		for i := 0; i < len(allIndexStr); i++ {
			for j := 0; j < len(allIndexStr[i]); j++ {
				if len(allIndexStr) > 1 {
					fmt.Printf("%s:", files[i])
				}
				if flags.flagN {
					fmt.Printf("%d:", allIndexStr[i][j]+1)
				}
				if len(allScanStr[i]) > allIndexStr[i][j] { // чтобы не выходить за пределы массива строк
					fmt.Println(allScanStr[i][allIndexStr[i][j]])
				}
			}
			if len(allIndexStr) > 1 && i != len(allIndexStr)-1 && (flags.flagA > 0 || flags.flagB > 0 || flags.flagC > 0) {
				fmt.Printf("--\n")
			}
		}
	}
}

func printAllIndexStr(allIndexStr [][]int) {
	for i := 0; i < len(allIndexStr); i++ {
		for j := 0; j < len(allIndexStr[i]); j++ {
			fmt.Println(allIndexStr[i][j])
		}
		fmt.Println()
	}
}

func printAllScanStr(allScanStr [][]string) {
	for i := 0; i < len(allScanStr); i++ {
		for j := 0; j < len(allScanStr[i]); j++ {
			fmt.Println(allScanStr[i][j])
		}
		fmt.Println()
	}
}
