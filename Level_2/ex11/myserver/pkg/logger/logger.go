package logger

import (
	"log"
	"net/http"
	"time"
)

func LogRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логирование информации о запросе
		log.Printf("[%s] %s %s", time.Now().Format(time.RFC3339), r.Method, r.URL.String())

		// Вызов следующего обработчика
		next(w, r)
	}
}
