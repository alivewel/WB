package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func flagh(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		return numericLessWithSuffix(lines[i], lines[j])
	})
	return lines
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

func main() {
	input := []string{"file2.txt.10M", "file1.txt.2K", "file3.txt.1G"}
	res := flagh(input)
	fmt.Println(res)
}
