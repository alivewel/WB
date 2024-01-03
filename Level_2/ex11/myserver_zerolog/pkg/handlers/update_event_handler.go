package handlers

import (
	"io"
	"myserver/pkg/event"

	"myserver/pkg/memorycache"

	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func UpdateEventHandler(cache *memorycache.Cache, logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logger.Warn().Msg("Метод не поддерживается")
			response := Response{Error: "Метод не поддерживается"}
			SendJSONResponse(w, http.StatusInternalServerError, response, logger)
			return
		} else {
			// Обработка запроса на создание события
			var eventData event.Event
			err := json.NewDecoder(r.Body).Decode(&eventData)
			if err != nil {
				logger.Warn().Msg("Ошибка разбора JSON")
				response := Response{Error: "Ошибка разбора JSON"}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
				return
			}

			err = cache.UpdateEvent(eventData, 5*time.Minute)
			if err != nil {
				logger.Warn().Msgf("Событие в кэше обновить не удалось: %v", err)
				response := Response{Error: err.Error()}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
				return
			}

			// Преобразование тела запроса в строку
			body, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Warn().Msgf("Ошибка чтения тела запроса")
				response := Response{Error: "Ошибка чтения тела запроса"}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
				return
			}
			requestData := string(body)

			logger.Info().Msg("Событие успешно обновлено!")

			response := Response{Result: requestData}
			SendJSONResponse(w, http.StatusOK, response, logger)
		}
	}
}
