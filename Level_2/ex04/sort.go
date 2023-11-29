package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func main() {
	lines := []string{"пятак", "пятка", "тяпка", "тяпка", "листок", "слиток", "столик", "кошка"}
	sortLines := make(map[string][]string, len(lines))
	// Выводим отсортированные руны
	for _, line := range lines {
		lowerLine := strings.ToLower(line)
		sortLine := sortLine(line) // Массив должен быть отсортирован по возрастанию
		sortLines[sortLine] = append(sortLines[sortLine], lowerLine)
	}


	sortLines3 := make(map[string][]string)
	for _, v := range sortLines {
		if len(v) > 1 { // Множества из одного элемента не должны попасть в результат
			firstValue := v[0] // Ключ - первое встретившееся в словаре слово из множества
			prevLine := ""
			for i := 0; i < len(v); i++ {
				if prevLine != v[i] { // В результате каждое слово должно встречаться только один раз
					sortLines3[firstValue] = append(sortLines3[firstValue], v[i])
					prevLine = v[i]
				}
			}
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
