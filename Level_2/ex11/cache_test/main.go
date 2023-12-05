package main

import (
	"cache_test/pkg/event"
	"cache_test/pkg/memorycache"

	// "errors"
	"fmt"
	"time"
)

// type Event struct {
// 	Summary string    `json:"summary"`
// 	Date    time.Time `json:"date"`
// }

func main() {
	// создаем событие и дату
	summary := "Мое событие"

	// дата приходит в таком формате date=2019-09-09
	dateString := "2019-09-09"

	eventInstance, err := event.NewEvent(summary, dateString)
	if err != nil {
		fmt.Println("Ошибка при создании мероприятия:", err)
		return
	}

	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	keyCache := summary + "_" + dateString

	cache.Set(keyCache, eventInstance, 5*time.Minute)

	// addCacheDay(cache)
	// addCacheWeek(cache)
	addCacheMonth(cache)

	// cache.PrintAll()
	// cache.PrintDay(12)
	// cache.PrintWeek(5)
	cache.PrintMonth(12)
}

// // функция которая создает новую структуру Event (конструктор)
// func NewEvent(summary, date string) (Event, error) {
// 	parsedTime, err := time.Parse("2006-01-02", date)
// 	if err != nil {
// 		return Event{}, fmt.Errorf("ошибка при разборе даты: %v", err)
// 	}

// 	if summary == "" || parsedTime == (time.Time{}) {
// 		return Event{}, errors.New("некорректные данные для создания события")
// 	}

// 	return Event{
// 		Summary: summary,
// 		Date:    parsedTime,
// 	}, nil
// }

func addCacheMonth(c *memorycache.Cache) {
	for i := 1; i <= 12; i++ {
		month := fmt.Sprintf("%02d", i)
		date := "2019-" + month + "-20"
		summary := "Мое событие " + month

		eventInstance, err := event.NewEvent(summary, date)
		if err != nil {
			fmt.Println("Ошибка при создании мероприятия:", err)
			return
		}

		keyCache := getKeyCache(summary, date)

		c.Set(keyCache, eventInstance, 5*time.Minute)
	}
}

func addCacheWeek(c *memorycache.Cache) {
	startDateStr := "2021-01-01"                            // начальная дата
	parsedTime, _ := time.Parse("2006-01-02", startDateStr) // переводим в формат time.Time
	// создаем дату с 1 по 12 неделю
	for i := 1; i <= 12; i++ {
		week := fmt.Sprintf("%02d", i)
		date := parsedTime.AddDate(0, 0, (i-1)*7)
		dateString := date.Format("2006-01-02") // перевод в строку

		summary := "Мое событие " + week

		eventInstance, err := event.NewEvent(summary, dateString)
		if err != nil {
			fmt.Println("Ошибка при создании мероприятия:", err)
			return
		}

		keyCache := getKeyCache(summary, dateString)

		c.Set(keyCache, eventInstance, 5*time.Minute)
	}
	// return nil
}

func addCacheDay(c *memorycache.Cache) {
	for i := 11; i <= 22; i++ {
		day := fmt.Sprintf("%02d", i)
		date := "2021-" + "01-" + day
		summary := "Мое событие " + day

		eventInstance, err := event.NewEvent(summary, date)
		if err != nil {
			fmt.Println("Ошибка при создании мероприятия:", err)
			return
		}

		keyCache := getKeyCache(summary, date)

		c.Set(keyCache, eventInstance, 5*time.Minute)
	}
}

func getKeyCache(summary, date string) string {
	if summary != "" && date != "" {
		return summary + "_" + date
	}
	return ""
}
