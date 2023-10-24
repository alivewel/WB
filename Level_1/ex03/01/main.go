package main

import "fmt"

func calculateSquareSum(numbers []int, resultChan chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	resultChan <- sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numGoroutines := 2 // Количество горутин для вычислений

	resultChan := make(chan int)

	chunkSize := len(numbers) / numGoroutines
	fmt.Printf("chunkSize: %d\n", chunkSize)

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		// fmt.Printf("%d: start %d end %d\n", i, start, end)
		if i == numGoroutines-1 {
			end = len(numbers)
		}
		fmt.Printf("%d: start %d end %d\n", i, start, end)
		fmt.Println(i, numbers[start:end])
		go calculateSquareSum(numbers[start:end], resultChan)
	}

	totalSum := 0

	for i := 0; i < numGoroutines; i++ {
		totalSum += <-resultChan
	}

	close(resultChan)

	fmt.Printf("Сумма квадратов: %d\n", totalSum)
}
