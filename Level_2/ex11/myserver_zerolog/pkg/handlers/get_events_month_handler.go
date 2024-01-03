package handlers

import (
	"myserver/pkg/memorycache"
	"strconv"

	"net/http"

	"github.com/rs/zerolog"
)

// curl "http://localhost:8080/events_for_month?month=5"
// Обработка запроса на получение с выбранного месяца
func GetEventsMonthHandler(cache *memorycache.Cache, logger zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		month := r.URL.Query().Get("month")

		monthInt, err := strconv.Atoi(month)
		if err != nil {
			response := Response{Error: "Ошибка преобразования строки в число"}
			SendJSONResponse(w, http.StatusBadRequest, response, logger)
			return
		}

		selectMonth, err := cache.GetFilterEventsByMonth(monthInt)
		if err != nil {
			response := Response{Error: "Ошибка парсинга URL запроса"}
			SendJSONResponse(w, http.StatusServiceUnavailable, response, logger)
			return
		}

		response := Response{Result: selectMonth}
		SendJSONResponse(w, http.StatusOK, response, logger)
	}
}
