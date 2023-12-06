package main

import (
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Создаем буфер для чтения данных
	buffer := make([]byte, 1024)

	for {
		// Читаем данные из соединения
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return
		}

		// Выводим полученные данные
		fmt.Printf("Получено: %s", buffer[:n])

		// Отправляем ответ обратно клиенту
		response := []byte("Сообщение получено\n")
		conn.Write(response)
		fmt.Printf("%s", response)
	}
}

func main() {
	// Устанавливаем адрес и порт
	address := "localhost"
	port := "8080"
	listenAddr := fmt.Sprintf("%s:%s", address, port)

	// Слушаем порт
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Ошибка прослушивания порта:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Сервер слушает на %s\n", listenAddr)

	for {
		// Принимаем входящее соединение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}

		// Запускаем обработку соединения в отдельной горутине
		go handleConnection(conn)
	}
}
