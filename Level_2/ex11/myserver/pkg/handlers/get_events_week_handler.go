package handlers

import (
	"myserver/pkg/memorycache"
	"strconv"

	"net/http"
)

// curl "http://localhost:8080/events_for_week?week=5"
// Обработка запроса на получение с выбранной недели
func GetEventsWeekHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		week := r.URL.Query().Get("week")

		weekInt, err := strconv.Atoi(week)
		if err != nil {
			response := Response{Error: "Ошибка преобразования строки в число"}
			SendJSONResponse(w, http.StatusBadRequest, response)
			return
		}

		selectWeek, err := cache.GetFilterEventsByDay(weekInt)
		if err != nil {
			response := Response{Error: "Ошибка парсинга URL запроса"}
			SendJSONResponse(w, http.StatusServiceUnavailable, response)
			return
		}

		response := Response{Result: selectWeek}
		SendJSONResponse(w, http.StatusOK, response)
	}
}
