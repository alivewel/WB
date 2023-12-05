package main

import (
	"myserver/pkg/event"
	"myserver/pkg/memorycache"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Event представляет структуру данных для хранения информации о мероприятии.
type Event struct {
	Summary string    `json:"summary"`
	Date    time.Time `json:"date"`
}

var (
	events      []Event
	eventsMutex sync.Mutex
)

func main() {
	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)
	
	http.HandleFunc("/create-event", logRequestMiddleware(createEventHandler))
	http.HandleFunc("/get-events", logRequestMiddleware(getEventsHandler))

	// Указываем порт для прослушивания
	port := 8080
	fmt.Printf("Сервер запущен на порту %d...\n", port)

	// Запуск сервера
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Обработка запроса на создание события
		var eventData Event
		err := json.NewDecoder(r.Body).Decode(&eventData)
		if err != nil {
			http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
			return
		}

		eventsMutex.Lock()
		defer eventsMutex.Unlock()

		// Добавление события в локальный список
		events = append(events, eventData)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Событие успешно добавлено.")
	} else {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}

// curl -X POST -H "Content-Type: application/json" -d '{"summary":"Мое событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/create-event

func getEventsHandler(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса на получение всех событий
	eventsMutex.Lock()
	defer eventsMutex.Unlock()

	// Конвертация списка событий в JSON
	eventsJSON, err := json.Marshal(events)
	if err != nil {
		http.Error(w, "Ошибка маршалинга событий в JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(eventsJSON)
}

// curl http://localhost:8080/get-events
// curl "http://localhost:8080/get-events?summary=Мое%20событие"

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логирование информации о запросе
		log.Printf("[%s] %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.String())

		// Вызов следующего обработчика
		next(w, r)
	}
}
