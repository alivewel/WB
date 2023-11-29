package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	lines := []string{"пятак", "пятка", "тяпка"}
	sortLines := make(map[string][]string, len(lines))
	// Выводим отсортированные руны
	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		sortLine := sortLine(line)
		// fmt.Printf("%s: %s\n", line, sortLine)
		sortLines[sortLine] = append(sortLines[sortLine], lowerLine)
	}

	for k, v := range sortLines {
		fmt.Printf("%s: %v\n", k, v)
	}
	// sortLines2 := make(map[string]string, len(lines))
	sortLines2 := make(map[string][]string)
	for k, v := range sortLines {
		firstValue := v[0]
		sortLines2[firstValue] = append(sortLines2[firstValue], k)
	}
	for k, v := range sortLines2 {
		fmt.Printf("%s: %v\n", k, v)
	}

	// sortLines3 := make(map[string]string, len(lines))
	sortLines3 := make(map[string][]string)
	for _, v := range sortLines {
		firstValue := v[0]
		for i := 0; i < len(v); i++ {
			sortLines3[firstValue] = append(sortLines3[firstValue], v[i])
		}
	}
	for k, v := range sortLines3 {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func sortLine(line string) string {
	// Разбиваем строку на руны
	runes := []rune{}
	for i := 0; i < len(line); {
		r, size := utf8.DecodeRuneInString(line[i:])
		runes = append(runes, r)
		i += size
	}
	// Сортируем массив рун
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
