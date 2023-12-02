package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <host> <port>")
		return
	}

	host := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to", conn.RemoteAddr())

	// Инициализация канала для обработки сигналов SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	wg.Add(2)

	go func() {
		defer wg.Done()
		readFromServer(conn)
	}()

	go func() {
		defer wg.Done()
		sendToServer(conn)
	}()

	// Ждем сигнала SIGTERM
	<-sigChan

	fmt.Println("\nReceived SIGTERM. Closing connection...")
	conn.Close()

	// Ожидаем завершения чтения и отправки
	wg.Wait()

	fmt.Println("Program terminated.")
}

func readFromServer(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed.")
			return
		}
		fmt.Print("Message from server:", message)
	}
}

func sendToServer(conn net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text to send to the server: \n")
		text, _ := reader.ReadString('\n')

		// Отправляем текст на сервер
		_, err := conn.Write([]byte(strings.TrimRight(text, "\n")))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
	}
}
