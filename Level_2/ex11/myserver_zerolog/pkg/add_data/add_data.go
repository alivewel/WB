package add_data

import (
	"myserver/pkg/event"
	"myserver/pkg/memorycache"

	"fmt"
	"time"
)

func AddCacheMonth(c *memorycache.Cache) {
	for i := 1; i <= 12; i++ {
		month := fmt.Sprintf("%02d", i)
		date := "2019-" + month + "-20"
		summary := "Мое событие " + month

		eventInstance, err := event.NewEvent(summary, date)
		if err != nil {
			fmt.Println("Ошибка при создании мероприятия:", err)
			return
		}

		keyCache := memorycache.GetKeyCache(eventInstance)

		c.Set(keyCache, eventInstance, 5*time.Minute)
	}
}

func AddCacheWeek(c *memorycache.Cache) {
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

		keyCache := memorycache.GetKeyCache(eventInstance)

		c.Set(keyCache, eventInstance, 5*time.Minute)
	}
}

func AddCacheDay(c *memorycache.Cache) {
	for i := 11; i <= 22; i++ {
		day := fmt.Sprintf("%02d", i)
		date := "2021-" + "01-" + day
		summary := "Мое событие " + day

		eventInstance, err := event.NewEvent(summary, date)
		if err != nil {
			fmt.Println("Ошибка при создании мероприятия:", err)
			return
		}

		keyCache := memorycache.GetKeyCache(eventInstance)

		c.Set(keyCache, eventInstance, 5*time.Minute)
	}
}
