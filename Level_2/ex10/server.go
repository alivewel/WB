package main

import (
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on :8081...")
	http.ListenAndServe(":8081", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	// 	return
	// }

	// Читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Преобразуем байтовый массив в строку
	message := string(body)

	fmt.Println("Received message:", message)

	response := "Message received successfully\n"
	io.WriteString(w, response)
}

// func handleRequest(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var message Message
// 	decoder := json.NewDecoder(r.Body)
// 	err := decoder.Decode(&message)
// 	if err != nil {
// 		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Println("Received message:", message.Text)

// 	response := "Message received successfully\n"
// 	io.WriteString(w, response)
// }

