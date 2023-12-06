package main

import (
	"io"
	"myserver/pkg/event"
	"myserver/pkg/handlers"
	"myserver/pkg/memorycache"
	"strconv"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// POST /create_event
// POST /update_event
// POST /delete_event
// GET /events_for_day
// GET /events_for_week
// GET /events_for_month

func main() {
	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	http.HandleFunc("/create-event", logRequestMiddleware(СreateEventHandler(cache)))
	http.HandleFunc("/events_for_day", logRequestMiddleware(GetEventsDayHandler(cache)))
	http.HandleFunc("/events_for_week", logRequestMiddleware(GetEventsWeekHandler(cache)))
	http.HandleFunc("/events_for_month", logRequestMiddleware(GetEventsMonthHandler(cache)))

	// Указываем порт для прослушивания
	port := 8080
	fmt.Printf("Сервер запущен на порту %d...\n", port)

	// Запуск сервера
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// curl -X POST -H "Content-Type: application/json" -d '{"summary":"Мое событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/create-event
func СreateEventHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Метод не поддерживается")
			response := handlers.Response{Error: "Метод не поддерживается"}
			handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
		} else {
			// Обработка запроса на создание события
			var eventData event.Event
			err := json.NewDecoder(r.Body).Decode(&eventData)
			if err != nil {
				log.Println("Ошибка разбора JSON")
				response := handlers.Response{Error: "Ошибка разбора JSON"}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}

			cache.AddEvent(eventData, 5*time.Minute)

			// Преобразование тела запроса в строку
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Ошибка чтения тела запроса")
				response := handlers.Response{Error: "Ошибка чтения тела запроса"}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}
			requestData := string(body)

			log.Printf("Событие успешно добавлено: %+v", eventData)

			response := handlers.Response{Result: requestData}
			handlers.SendJSONResponse(w, http.StatusOK, response)
		}
	}
}

func UpdateEventHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Метод не поддерживается")
			response := handlers.Response{Error: "Метод не поддерживается"}
			handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
		} else {
			// Обработка запроса на создание события
			var eventData event.Event
			err := json.NewDecoder(r.Body).Decode(&eventData)
			if err != nil {
				log.Println("Ошибка разбора JSON")
				response := handlers.Response{Error: "Ошибка разбора JSON"}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}

			err = cache.UpdateEvent(eventData, 5*time.Minute)
			if err != nil {
				log.Println(err.Error())
				response := handlers.Response{Error: err.Error()}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}

			// Преобразование тела запроса в строку
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Ошибка чтения тела запроса")
				response := handlers.Response{Error: "Ошибка чтения тела запроса"}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}
			requestData := string(body)

			log.Printf("Событие успешно обновлено: %+v", eventData)

			response := handlers.Response{Result: requestData}
			handlers.SendJSONResponse(w, http.StatusOK, response)
		}
	}
}

func DeleteEventHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Метод не поддерживается")
			response := handlers.Response{Error: "Метод не поддерживается"}
			handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
		} else {
			// Обработка запроса на создание события
			var eventData event.Event
			err := json.NewDecoder(r.Body).Decode(&eventData)
			if err != nil {
				log.Println("Ошибка разбора JSON")
				response := handlers.Response{Error: "Ошибка разбора JSON"}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}

			err = cache.DeleteEvent(eventData)
			if err != nil {
				log.Println(err.Error())
				response := handlers.Response{Error: err.Error()}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}

			// Преобразование тела запроса в строку
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Ошибка чтения тела запроса")
				response := handlers.Response{Error: "Ошибка чтения тела запроса"}
				handlers.SendJSONResponse(w, http.StatusInternalServerError, response)
			}
			requestData := string(body)

			log.Printf("Событие успешно удалено: %+v", eventData)

			response := handlers.Response{Result: requestData}
			handlers.SendJSONResponse(w, http.StatusOK, response)
		}
	}
}


// curl "http://localhost:8080/events_for_day?day=12"
// Обработка запроса на получение с выбранного дня
func GetEventsDayHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		day := r.URL.Query().Get("day")

		dayInt, err := strconv.Atoi(day)
		if err != nil {
			response := handlers.Response{Error: "Ошибка преобразования строки в число"}
			handlers.SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		selectDay, err := cache.GetFilterEventsByDay(dayInt)
		if err != nil {
			response := handlers.Response{Error: "Ошибка парсинга URL запроса"}
			handlers.SendJSONResponse(w, http.StatusServiceUnavailable, response)
			return
		}

		response := handlers.Response{Result: selectDay}
		handlers.SendJSONResponse(w, http.StatusOK, response)
	}
}

// curl "http://localhost:8080/events_for_week?week=5"
// Обработка запроса на получение с выбранной недели
func GetEventsWeekHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		week := r.URL.Query().Get("week")

		weekInt, err := strconv.Atoi(week)
		if err != nil {
			response := handlers.Response{Error: "Ошибка преобразования строки в число"}
			handlers.SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		selectWeek, err := cache.GetFilterEventsByDay(weekInt)
		if err != nil {
			response := handlers.Response{Error: "Ошибка парсинга URL запроса"}
			handlers.SendJSONResponse(w, http.StatusServiceUnavailable, response)
			return
		}

		response := handlers.Response{Result: selectWeek}
		handlers.SendJSONResponse(w, http.StatusOK, response)
	}
}

// curl "http://localhost:8080/events_for_month?month=5"
// Обработка запроса на получение с выбранного месяца
func GetEventsMonthHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		month := r.URL.Query().Get("month")

		monthInt, err := strconv.Atoi(month)
		if err != nil {
			response := handlers.Response{Error: "Ошибка преобразования строки в число"}
			handlers.SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		selectMonth, err := cache.GetFilterEventsByDay(monthInt)
		if err != nil {
			response := handlers.Response{Error: "Ошибка парсинга URL запроса"}
			handlers.SendJSONResponse(w, http.StatusServiceUnavailable, response)
			return
		}

		response := handlers.Response{Result: selectMonth}
		handlers.SendJSONResponse(w, http.StatusOK, response)
	}
}

// curl http://localhost:8080/get-events
// curl "http://localhost:8080/get-events?summary=Мое%20событие"
// curl "http://localhost:8080/get-events?day=12"

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
