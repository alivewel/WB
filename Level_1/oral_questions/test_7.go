package main

import (
	"fmt"
	"syscall"
)

func main() {
	filename := "/tmp/testfile.txt"
	flags := syscall.O_RDONLY // открыть файл в режиме только для чтения
	mode  := uint32 ( 0666 ) // права доступа к файлу - rw-rw-rw-
	fd, err := syscall.Open(filename, flags, mode)
	if err != nil {
		fmt.Printf("Ошибка открытия файла %s: %s\n", filename, err)
		return
	}
	defer syscall.Close(fd)
	fmt.Printf("Файл %s открыт успешно!\n", filename)
}
