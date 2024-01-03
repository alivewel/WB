package handlers

import (
	"myserver/pkg/memorycache"
	"strconv"

	"net/http"

	"github.com/rs/zerolog"
)

// curl "http://localhost:8080/events_for_day?day=12"
// Обработка запроса на получение с выбранного дня
func GetEventsDayHandler(cache *memorycache.Cache, logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		day := r.URL.Query().Get("day")

		dayInt, err := strconv.Atoi(day)
		if err != nil {
			response := Response{Error: "Ошибка преобразования строки в число"}
			SendJSONResponse(w, http.StatusBadRequest, response, logger)
			return
		}

		selectDay, err := cache.GetFilterEventsByDay(dayInt)
		if err != nil {
			response := Response{Error: "Ошибка парсинга URL запроса"}
			SendJSONResponse(w, http.StatusServiceUnavailable, response, logger)
			return
		}

		response := Response{Result: selectDay}
		SendJSONResponse(w, http.StatusOK, response, logger)
	}
}
