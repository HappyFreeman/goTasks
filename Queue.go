package main

import "fmt"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}

type Queue [T any] struct {
    elements []T
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue[T]) Enqueue(value T) {
    q.elements = append(q.elements, value)
}

// Dequeue удаляет элемент из начала очереди и возвращает его
func (q *Queue[T]) Dequeue() (T, bool) {
	var zeroValue T
    if len(q.elements) == 0 {
        return zeroValue, false
    }
    element := q.elements[0]
    q.elements = q.elements[1:]
    return element, true
}

// Peek возвращает элемент из начала очереди
func (q *Queue[T]) Peek() (T, bool) {
	var zeroValue T
    if len(q.elements) == 0 {
        return zeroValue, false
    }
    element := q.elements[0]
    return element, true
}

func main() {
    queue := &Queue[int]{}
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    fmt.Println(queue.Dequeue()) // Вывод: 1, true
    fmt.Println(queue.Dequeue()) // Вывод: 2, true
    fmt.Println(queue.Dequeue()) // Вывод: 3, true
    fmt.Println(queue.Dequeue()) // Вывод: 0, false
}
