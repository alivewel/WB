package handlers

import (
	"myserver/pkg/memorycache"
	"strconv"

	"net/http"
)

// curl "http://localhost:8080/events_for_month?month=5"
// Обработка запроса на получение с выбранного месяца
func GetEventsMonthHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		month := r.URL.Query().Get("month")

		monthInt, err := strconv.Atoi(month)
		if err != nil {
			response := Response{Error: "Ошибка преобразования строки в число"}
			SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		selectMonth, err := cache.GetFilterEventsByDay(monthInt)
		if err != nil {
			response := Response{Error: "Ошибка парсинга URL запроса"}
			SendJSONResponse(w, http.StatusServiceUnavailable, response)
			return
		}

		response := Response{Result: selectMonth}
		SendJSONResponse(w, http.StatusOK, response)
	}
}
