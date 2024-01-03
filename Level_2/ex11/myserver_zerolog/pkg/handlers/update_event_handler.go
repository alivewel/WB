package handlers

import (
	"io"
	"myserver/pkg/event"

	"myserver/pkg/memorycache"

	"encoding/json"
	"log"
	"net/http"
	"time"
)

func UpdateEventHandler(cache *memorycache.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			log.Println("Метод не поддерживается")
			response := Response{Error: "Метод не поддерживается"}
			SendJSONResponse(w, http.StatusInternalServerError, response)
		} else {
			// Обработка запроса на создание события
			var eventData event.Event
			err := json.NewDecoder(r.Body).Decode(&eventData)
			if err != nil {
				log.Println("Ошибка разбора JSON")
				response := Response{Error: "Ошибка разбора JSON"}
				SendJSONResponse(w, http.StatusInternalServerError, response)
			}

			err = cache.UpdateEvent(eventData, 5*time.Minute)
			if err != nil {
				log.Println(err.Error())
				response := Response{Error: err.Error()}
				SendJSONResponse(w, http.StatusInternalServerError, response)
			}

			// Преобразование тела запроса в строку
			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Ошибка чтения тела запроса")
				response := Response{Error: "Ошибка чтения тела запроса"}
				SendJSONResponse(w, http.StatusInternalServerError, response)
			}
			requestData := string(body)

			log.Printf("Событие успешно обновлено: %+v", eventData)

			response := Response{Result: requestData}
			SendJSONResponse(w, http.StatusOK, response)
		}
	}
}
