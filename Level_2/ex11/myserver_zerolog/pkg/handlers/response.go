package handlers

import (
	"encoding/json"
	"net/http"
	
	"github.com/rs/zerolog"
)

// Response структура для JSON-ответа
type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// SendJSONResponse отправляет JSON-ответ с заданным кодом состояния
func SendJSONResponse(w http.ResponseWriter, statusCode int, response Response, logger zerolog.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, err := json.Marshal(response)
	if err != nil {
		// log.Printf("Ошибка при маршалинге JSON: %v", err)
		logger.Warn().Msgf("Ошибка при маршалинге JSON: %v", err)
		// В случае ошибки маршалинга, возвращаем HTTP 500 и логируем ошибку
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
