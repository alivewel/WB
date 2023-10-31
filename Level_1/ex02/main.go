package main

import "fmt"

func squares(c chan int, numbers []int) {
	for _, num := range numbers {
		c <- num * num
	}
	close(c)
}

func main() {
	arrayInt := []int{2, 4, 6, 8, 10}

	c := make(chan int)

	go squares(c, arrayInt)

	for val := range c {
		fmt.Println(val)
	}
}

// Когда мы доходим до строки for val := range c, канал c уже закрыт.
// Это обусловлено строкой close(c) в функции squares. Когда канал закрыт,
// цикл for val := range c будет продолжаться до тех пор, пока не будут
//  прочитаны все значения из канала, после чего цикл завершит выполнение.

// Закрытие канала важно, чтобы горутина, которая отправляет значения в канал
//  (в данном случае функция squares), могла сигнализировать, что она больше не будет
//  отправлять новые значения. Это помогает избежать блокировки и завершить цикл
//  чтения данных из канала в главной горутине, когда все данные уже получены.

// Канал, созданный с помощью make(chan int), будет безразмерным (unbuffered),
// что означает, что он может принимать и отправлять значения только тогда,
// когда обе операции, отправка и прием, готовы. Это означает, что при отправке значения
// в безразмерный канал (c <- value), отправящая горутина будет заблокирована до тех пор,
// пока другая горутина не считает это значение из канала (<-c).
