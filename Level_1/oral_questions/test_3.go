package main

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: []int{},
	}
}

func (q *Queue) Enqueue(value int) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Очередь пуста")
	}

	value := q.data[0]
	q.data = q.data[1:]
	return value, nil
}

func main() {
	queue := NewQueue()

	// Добавляем элементы в очередь
	for i := 1; i <= 5; i++ {
		queue.Enqueue(i)
	}

	// Извлекаем элементы из очереди
	for i := 0; i < 5; i++ {
		value, err := queue.Dequeue()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Извлечено: %d\n", value)
		}
	}
}
