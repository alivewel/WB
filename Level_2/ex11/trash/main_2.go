package main

import (
	"errors"
	"fmt"
	"time"
)

type Event struct {
	Summary string    `json:"summary"`
	Date    time.Time `json:"date"`
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case Event:
		fmt.Printf("%q is Event\n", v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)

	// создаем событие и дату
	summary := "Мое событие"

	// дата приходит в таком формате date=2019-09-09
	dateString := "2019-09-09"

	eventInstance, err := NewEvent(summary, dateString)
	if err != nil {
		fmt.Println("Ошибка при создании мероприятия:", err)
		return
	}
	do(eventInstance)
}

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
