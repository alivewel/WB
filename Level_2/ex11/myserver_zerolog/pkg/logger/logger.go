package logger

import (
	"net/http"

	"github.com/rs/zerolog"
)

func LogRequestMiddleware(next http.HandlerFunc, logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Логирование информации о запросе
		logger.Info().Msgf("Method: %s, URL: %s", r.Method, r.URL.String())

		// Вызов следующего обработчика
		next(w, r)
	}
}
