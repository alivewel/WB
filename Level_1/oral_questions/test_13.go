package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriterSize(file, 5) // Маленький буфер размером 5 байт

	data := []byte("Длинные данные для записи в файл")
	_, err = writer.Write(data)
	if err != nil {
		panic(err)
	}

	err = writer.Flush() // Попытка записать оставшиеся данные из буфера
	if err != nil {
		panic(err)
	}
}
