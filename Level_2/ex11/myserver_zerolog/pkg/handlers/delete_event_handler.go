package handlers

import (
	"io"
	"myserver/pkg/event"
	"myserver/pkg/memorycache"

	"encoding/json"
	"net/http"
	"github.com/rs/zerolog"
)

func DeleteEventHandler(cache *memorycache.Cache, logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logger.Warn().Msg("Метод не поддерживается")
			response := Response{Error: "Метод не поддерживается"}
			SendJSONResponse(w, http.StatusInternalServerError, response, logger)
		} else {
			// Обработка запроса на создание события
			var eventData event.Event
			err := json.NewDecoder(r.Body).Decode(&eventData)
			if err != nil {
				logger.Warn().Msg("Ошибка разбора JSON")
				response := Response{Error: "Ошибка разбора JSON"}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
			}

			err = cache.DeleteEvent(eventData)
			if err != nil {
				logger.Warn().Msgf("Событие в кэше удалить не удалось: %v", err)
				response := Response{Error: err.Error()}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
			}

			// Преобразование тела запроса в строку
			body, err := io.ReadAll(r.Body)
			if err != nil {
				logger.Warn().Msgf("Ошибка чтения тела запроса")
				response := Response{Error: "Ошибка чтения тела запроса"}
				SendJSONResponse(w, http.StatusInternalServerError, response, logger)
			}
			requestData := string(body)

			logger.Info().Msg("Событие успешно добавлено!")

			response := Response{Result: requestData}
			SendJSONResponse(w, http.StatusOK, response, logger)
		}
	}
}
