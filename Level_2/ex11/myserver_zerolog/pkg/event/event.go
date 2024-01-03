package event

import (
	"errors"
	"fmt"
	"time"
)

// временное дублирование
// потом вынести в отдельный файл
type Event struct {
	Summary string    `json:"summary"`
	Date    time.Time `json:"date"`
}

// функция которая создает новую структуру Event (конструктор)
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
