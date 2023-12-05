package main

import (
	"fmt"
	"time"
)

func generateDatesForWeekNumber(year, weekNumber int) []string {
	// Создаем временную метку для первого дня года и недели
	startDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, (weekNumber-1)*7)

	// Создаем срез для хранения дат
	dates := make([]string, 7)

	// Генерируем даты для каждого дня недели
	for i := 0; i < 7; i++ {
		currentDate := startDate.AddDate(0, 0, i)
		fmt.Println(currentDate)
		dates[i] = currentDate.Format("2006-01-02")
	}

	return dates
}

func main() {
	year := 2019
	weekNumber := 10

	// Генерируем даты для всех дней недели с номером 10 в 2019 году
	dates := generateDatesForWeekNumber(year, weekNumber)

	// Выводим результат
	for _, date := range dates {
		fmt.Println(date)
	}
}
