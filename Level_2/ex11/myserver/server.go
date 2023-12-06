package main

import (
	"myserver/pkg/event"
	"myserver/pkg/memorycache"
	"strconv"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	http.HandleFunc("/create-event", logRequestMiddleware(createEventHandler(cache)))
	http.HandleFunc("/get-events", logRequestMiddleware(getEventsHandler(cache)))

	// Указываем порт для прослушивания
	port := 8080
	fmt.Printf("Сервер запущен на порту %d...\n", port)

	// Запуск сервера
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func createEventHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			// Обработка запроса на создание события
			var eventData event.Event
			err := json.NewDecoder(r.Body).Decode(&eventData)
			if err != nil {
				http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
				return
			}

			cache.AddEvent(eventData, 5*time.Minute)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Событие успешно добавлено.")
		} else {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	}
}

// curl -X POST -H "Content-Type: application/json" -d '{"summary":"Мое событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/create-event

func getEventsHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Обработка запроса на получение всех событий
		summary := r.URL.Query().Get("summary")
		fmt.Println(summary)

		days := r.URL.Query().Get("days")
		fmt.Println(days)

		daysInt, err := strconv.Atoi(days)
		if err != nil {
			// Обработка ошибки, если преобразование не удалось
			http.Error(w, "Ошибка преобразования строки в число", http.StatusBadRequest)
			return
		}

		addCacheDay(cache)

		selectDay, err := cache.GetFilterEventsByDay(daysInt)
		if err != nil {
			http.Error(w, "Ошибка парсинга URL запроса", http.StatusBadRequest)
			return
		}
		fmt.Println(selectDay)
		// считать данные с доменного запроса
		// отправить запрос в БД
		// получить данные и вывести их на экран

		// cache.PrintAll()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// w.Write("Мир!")
	}
}

// curl http://localhost:8080/get-events
// curl "http://localhost:8080/get-events?summary=Мое%20событие"
// curl "http://localhost:8080/get-events?days=12"

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логирование информации о запросе
		log.Printf("[%s] %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.String())

		// Вызов следующего обработчика
		next(w, r)
	}
}

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

		keyCache := memorycache.GetKeyCache(eventInstance)

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

		keyCache := memorycache.GetKeyCache(eventInstance)

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

		keyCache := memorycache.GetKeyCache(eventInstance)

		c.Set(keyCache, eventInstance, 5*time.Minute)
	}
}
