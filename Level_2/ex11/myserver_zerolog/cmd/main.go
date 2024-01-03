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

	http.HandleFunc("/create_event", logger.LogRequestMiddleware(handlers.СreateEventHandler(cache, zero_logger), zero_logger))
	http.HandleFunc("/update_event", logger.LogRequestMiddleware(handlers.UpdateEventHandler(cache, zero_logger), zero_logger))
	http.HandleFunc("/delete_event", logger.LogRequestMiddleware(handlers.DeleteEventHandler(cache, zero_logger), zero_logger))
	http.HandleFunc("/events_for_day", logger.LogRequestMiddleware(handlers.GetEventsDayHandler(cache, zero_logger), zero_logger))
	http.HandleFunc("/events_for_week", logger.LogRequestMiddleware(handlers.GetEventsWeekHandler(cache, zero_logger), zero_logger))
	http.HandleFunc("/events_for_month", logger.LogRequestMiddleware(handlers.GetEventsMonthHandler(cache, zero_logger), zero_logger))

	// Указываем порт для прослушивания
	port := flag.Int("port", 8080, "Порт для прослушивания")
	flag.Parse()

	// fmt.Printf("Сервер запущен на порту %d...\n", *port)
	zero_logger.Info().Msgf("Сервер запущен на порту %d", *port)

	// Запуск сервера
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		zero_logger.Warn().Msgf("Ошибка запуска сервера:", err)
	}
}

// create-event
// curl -X POST -H "Content-Type: application/json" -d '{"summary":"Мое событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/create_event

// curl -X POST -H "Content-Type: application 2/json" -d '{"summary":"Мое событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/update_event
// curl -X POST -H "Content-Type: application 2/json" -d '{"summary":"Мое обновленное событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/update_event
