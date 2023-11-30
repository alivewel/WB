package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	FlagK int
	FlagN bool
	FlagR bool
	FlagU bool
	FlagM bool
	FlagB bool
	FlagC bool
	FlagH bool
}

var flags Flags

// тесты с функцией init не работают, пришлось создать свою функцию
func initFlags() {
	flag.IntVar(&flags.FlagK, "k", 0, "Specifying the column to sort")
	flag.BoolVar(&flags.FlagN, "n", false, "Sort by numeric value")
	flag.BoolVar(&flags.FlagR, "r", false, "Sort in reverse order")
	flag.BoolVar(&flags.FlagU, "u", false, "Do not print duplicate lines")
	flag.BoolVar(&flags.FlagM, "M", false, "Sort by month name")
	flag.BoolVar(&flags.FlagB, "b", false, "Ignore trailing spaces")
	flag.BoolVar(&flags.FlagC, "c", false, "Check if data is sorted")
	flag.BoolVar(&flags.FlagH, "h", false, "Sort by numeric value taking into account suffixes")
	flag.Parse()
}

func main() {
	initFlags()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: sort-go [OPTIONS] FILENAME [FILENAME...]")
		os.Exit(1)
	}

	fileName := args[0]

	lines, err := getLines(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sortedLines := customSort(lines)
	printLines(sortedLines)
}

func customSort(lines []string) []string {
	if areFlagsOff() {
		notFlags(lines)
	}
	if flags.FlagK > 0 {
		FlagK(lines)
	}
	if flags.FlagR {
		FlagR(lines)
	}
	if flags.FlagN {
		FlagN(lines)
	}
	if flags.FlagB {
		FlagB(lines)
	}
	if flags.FlagM {
		FlagM(lines)
	}
	if flags.FlagC && !FlagC(lines) {
		fmt.Printf("sort: -: disorder:\n")
	}
	if flags.FlagH {
		FlagH(lines)
	}
	return lines
}

func notFlags(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
	return lines
}

func FlagK(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		lineI := strings.Split(lines[i], " ")
		lineJ := strings.Split(lines[j], " ")
		if len(lineI) < flags.FlagK {
			if flags.FlagR {
				return true
			}
			return false
		}
		if len(lineJ) < flags.FlagK {
			if flags.FlagR {
				return false
			}
			return true
		}
		if flags.FlagR {
			return lineI[flags.FlagK-1] > lineJ[flags.FlagK-1]
		}
		return lineI[flags.FlagK-1] < lineJ[flags.FlagK-1]
	})
	return lines
}

func FlagR(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] > lines[j]
	})
	return lines
}

func FlagM(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return getMonthNumber(lines[i]) < getMonthNumber(lines[j])
	})
	return lines
}

func FlagB(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return strings.TrimSpace(lines[i]) < strings.TrimSpace(lines[j])
	})
	return lines
}

func FlagC(lines []string) bool {
	return sort.SliceIsSorted(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func numericLess(i, j string) bool {
	// Если обе строки числовые, сравниваем их как числа
	if isNumeric(i) && isNumeric(j) {
		numI, _ := strconv.Atoi(i)
		numJ, _ := strconv.Atoi(j)
		return numI < numJ
	}
	// В противном случае, сравниваем их лексикографически
	return i < j
}

// сортировка числовая, а не лексикографическая
func FlagN(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return numericLess(lines[i], lines[j])
	})
	return lines
}

func FlagH(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return numericLessWithSuffix(lines[i], lines[j])
	})
	return lines
}

func getLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var scanStr []string
	for scanner.Scan() {
		line := scanner.Text()
		scanStr = append(scanStr, line)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return scanStr, nil
}

func areFlagsOff() bool {
	return flags.FlagK == 0 &&
		!flags.FlagN &&
		!flags.FlagR &&
		!flags.FlagM &&
		!flags.FlagB &&
		!flags.FlagC &&
		!flags.FlagH
}

func getMonthNumber(monthName string) int {
	months := map[string]int{
		"январь":    1,
		"февраль":   2,
		"март":      3,
		"апрель":    4,
		"май":       5,
		"июнь":      6,
		"июль":      7,
		"август":    8,
		"сентябрь":  9,
		"октябрь":   10,
		"ноябрь":    11,
		"декабрь":   12,
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}

	monthName = strings.ToLower(monthName)

	number, found := months[monthName]
	if !found {
		return 0
	}

	return number
}

func numericLessWithSuffix(a, b string) bool {
	// Разделяем строку на основную часть и суффикс
	aMain, aSuffix := splitMainAndSuffix(a)
	bMain, bSuffix := splitMainAndSuffix(b)

	// Если основные части не равны, сравниваем их
	if aMain != bMain {
		return aMain < bMain
	}

	// Если основные части равны, преобразуем суффиксы в числа и сравниваем
	aNum, errA := strconv.Atoi(aSuffix)
	bNum, errB := strconv.Atoi(bSuffix)

	// Если преобразование прошло успешно, сравниваем числа
	if errA == nil && errB == nil {
		return aNum < bNum
	}

	// Иначе, если хотя бы один из суффиксов не является числом,
	// сравниваем строки как текст
	return a < b
}

func splitMainAndSuffix(s string) (main, suffix string) {
	parts := strings.Split(s, ".")
	if len(parts) > 1 {
		// Если строка содержит точку, считаем, что суффикс — это все, что после точки
		return parts[0], parts[1]
	}
	// Если точка отсутствует, считаем, что нет суффикса
	return s, ""
}

func printLines(lines []string) {
	if !flags.FlagC {
		prevLine := ""
		for _, line := range lines {
			if flags.FlagU {
				if prevLine != line {
					fmt.Printf("%s\n", line)
					prevLine = line
				}
			} else {
				fmt.Printf("%s\n", line)
			}
		}
	}
}
