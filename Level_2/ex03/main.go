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
	flagk int
	flagn bool
	flagr bool
	flagu bool
	flagM bool
	flagb bool
	flagc bool
	flagh bool
}

var flags Flags

// тесты с функцией init не работают, пришлось создать свою функцию
func initFlags() {
	flag.IntVar(&flags.flagk, "k", 0, "Specifying the column to sort")
	flag.BoolVar(&flags.flagn, "n", false, "Sort by numeric value")
	flag.BoolVar(&flags.flagr, "r", false, "Sort in reverse order")
	flag.BoolVar(&flags.flagu, "u", false, "Do not print duplicate lines")
	flag.BoolVar(&flags.flagM, "M", false, "Sort by month name")
	flag.BoolVar(&flags.flagb, "b", false, "Ignore trailing spaces")
	flag.BoolVar(&flags.flagc, "c", false, "Check if data is sorted")
	flag.BoolVar(&flags.flagh, "h", false, "Sort by numeric value taking into account suffixes")
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
	if flags.flagr {
		flagr(lines)
	}
	if flags.flagn {
		flagn(lines)
	}
	if flags.flagb {
		flagb(lines)
	}
	if flags.flagc && !flagc(lines) {
		fmt.Printf("sort: -: disorder:\n")
	}
	return lines
}

func notFlags(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
	return lines
}

func flagr(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] > lines[j]
	})
	return lines
}

func flagb(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return strings.TrimSpace(lines[i]) < strings.TrimSpace(lines[j])
	})
	return lines
}

func flagc(lines []string) bool {
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
func flagn(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return numericLess(lines[i], lines[j])
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
	// Проверяем значения по умолчанию для каждого флага
	return flags.flagk == 0 &&
		!flags.flagn &&
		!flags.flagr &&
		// !flags.flagu &&
		!flags.flagM &&
		!flags.flagb &&
		!flags.flagc &&
		!flags.flagh
}

func printLines(lines []string) {
	if !flags.flagc {
		prevLine := ""
		for _, line := range lines {
			if flags.flagu {
				if prevLine != line {
					fmt.Printf("%s\n", line)
					prevLine = line
				}
			} else {
				fmt.Printf("%s\n", line)
			}
		}
	}
	// else {

	// }
}
