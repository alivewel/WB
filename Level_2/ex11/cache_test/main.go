package main

import (
	"cache_test/pkg/memorycache"
	"errors"
	"fmt"
	"time"
)

type Event struct {
	Summary string    `json:"summary"`
	Date    time.Time `json:"date"`
}

func main() {
	// создаем событие и дату
	summary := "Мое событие"
	_ = summary
	// дата приходит в таком формате date=2019-09-09
	// dateString := "2023-12-31T23:59:59Z"
	dateString := "2019-09-09"
	parsedTime, err := time.Parse("2006-01-02", dateString) // проверка валидности даты
	if err != nil {
		fmt.Println("Ошибка при разборе даты:", err)
		return
	}
	fmt.Println(parsedTime)

	// в качестве ключа сделать название мероприятия_2023-12-31
	// в качестве значения структуру мероприятия

	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	keyCache := summary + "_" + dateString
	fmt.Println(keyCache)

	eventInstance := Event{
		Summary: summary,
		Date:    parsedTime,
	}

	cache.Set(keyCache, eventInstance, 5*time.Minute)

	//
	value, _ := cache.Get(keyCache)

	fmt.Println(value)
}

// функция которая создает новую структуру Event (конструктор)
// ПРОТЕСТИРОВАТЬ!!!
func NewEvent(summary, date string) (Event, error) {
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return Event{}, fmt.Errorf("ошибка при разборе даты: %v", err)
	}

	if summary == "" || parsedTime == (time.Time{}) {
		return Event{}, errors.New("некорректные данные для создания события")
	}

	return Event{
		Summary: summary,
		Date:    parsedTime,
	}, nil
}

// func NewEvent(summary, date string) Event {
// 	parsedTime, err := time.Parse("2006-01-02", date) // проверка валидности даты
// 	if err != nil {
// 		fmt.Println("Ошибка при разборе даты:", err)
// 		return Event{}
// 	}
// 	return Event{
// 		Summary: summary,
// 		Date:    parsedTime,
// 	}
// }
