package main

import (
	"flag"
	"myserver/pkg/handlers"
	"myserver/pkg/logger"
	"myserver/pkg/memorycache"

	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a container for the cache
	cache := memorycache.New(5*time.Minute, 10*time.Minute)

	http.HandleFunc("/create-event", logger.LogRequestMiddleware(handlers.СreateEventHandler(cache)))
	http.HandleFunc("/update_event", logger.LogRequestMiddleware(handlers.UpdateEventHandler(cache)))
	http.HandleFunc("/delete_event", logger.LogRequestMiddleware(handlers.DeleteEventHandler(cache)))
	http.HandleFunc("/events_for_day", logger.LogRequestMiddleware(handlers.GetEventsDayHandler(cache)))
	http.HandleFunc("/events_for_week", logger.LogRequestMiddleware(handlers.GetEventsWeekHandler(cache)))
	http.HandleFunc("/events_for_month", logger.LogRequestMiddleware(handlers.GetEventsMonthHandler(cache)))

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
