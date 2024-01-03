package main

import (
	"flag"
	"myserver/pkg/handlers"
	"myserver/pkg/logger"
	"myserver/pkg/memorycache"
	"os"

	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	// Создание логгера с выводом в стандартный вывод (stdout)
	zero_logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	http.HandleFunc("/create-event", logger.LogRequestMiddleware(handlers.СreateEventHandler(cache, zero_logger), zero_logger))
	// http.HandleFunc("/create-event", handlers.СreateEventHandler(cache, zero_logger))
	// http.HandleFunc("/update_event", logger.LogRequestMiddleware(handlers.UpdateEventHandler(cache)))
	// http.HandleFunc("/delete_event", logger.LogRequestMiddleware(handlers.DeleteEventHandler(cache)))
	// http.HandleFunc("/events_for_day", logger.LogRequestMiddleware(handlers.GetEventsDayHandler(cache)))
	// http.HandleFunc("/events_for_week", logger.LogRequestMiddleware(handlers.GetEventsWeekHandler(cache)))
	// http.HandleFunc("/events_for_month", logger.LogRequestMiddleware(handlers.GetEventsMonthHandler(cache)))

	// Указываем порт для прослушивания
	port := flag.Int("port", 8080, "Порт для прослушивания")
	flag.Parse()

	fmt.Printf("Сервер запущен на порту %d...\n", *port)

	// Запуск сервера
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}

// create-event
// curl -X POST -H "Content-Type: application/json" -d '{"summary":"Мое событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/create-event
