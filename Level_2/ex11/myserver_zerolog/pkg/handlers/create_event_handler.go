package handlers

import (
	"bytes"
	"io"

	"myserver/pkg/event"

	"myserver/pkg/memorycache"

	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

// curl -X POST -H "Content-Type: application/json" -d '{"summary":"Мое событие","date":"2023-12-31T23:59:59Z"}' http://localhost:8080/create_event
func СreateEventHandler(cache *memorycache.Cache, logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			messageError := "Метод не поддерживается"
			logger.Warn().Msg(messageError)
			response := Response{Error: messageError}
			SendJSONResponse(w, http.StatusInternalServerError, response, logger)
		} else {
			// Обработка запроса на создание события
			var eventData event.Event
			body, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Warn().Msg("Ошибка разбора JSON")
				response := Response{Error: "Ошибка разбора JSON"}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
				return
			}

			// Создаем новый читатель для r.Body
			newBody := io.NopCloser(bytes.NewBuffer(body))

			// Декодируем данные из нового читателя
			err = json.NewDecoder(newBody).Decode(&eventData)
			if err != nil {
				logger.Warn().Msgf("Ошибка разбора JSON: %v", err)
				response := Response{Error: "Ошибка разбора JSON"}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
				return
			}

			// Добавляем событие в кэш
			err = cache.AddEvent(eventData, 5*time.Minute)
			if err != nil {
				logger.Warn().Msgf("Событие в кэш добавить не удалось: %v", err)
				response := Response{Error: err.Error()}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
				return
			}

			// Преобразование тела запроса в строку
			requestData := string(body)

			logger.Info().Msg("Событие успешно добавлено!")
			response := Response{Result: requestData}
			SendJSONResponse(w, http.StatusOK, response, logger)
		}
	}
}
