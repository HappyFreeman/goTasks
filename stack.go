package main

import "fmt"

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}

type Stack [T any] struct {
    elements []T
}

// Push добавляет элемент в стек
func (s *Stack[T]) Push(value T) {
    s.elements = append(s.elements, value)
}


// Pop удаляет верхний элемент из стека и возвращает его
func (s *Stack[T]) Pop() (T, bool) {
    var zeroValue T
    if len(s.elements) == 0 {
        return zeroValue, false
    }
    index := len(s.elements) - 1
    element := s.elements[index]
    s.elements = s.elements[:index]
    return element, true
}

// Peek возвращает верхний элемент из стека
func (s *Stack[T]) Peek() (T, bool) {
	var zeroValue T
    if len(s.elements) == 0 {
        return zeroValue, false
    }
	index := len(s.elements) - 1
	element := s.elements[index]
	return element, true
}


func main() {
    stack := &Stack[int]{}
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)
    fmt.Println(stack.Pop()) // Вывод: 3, true
    fmt.Println(stack.Pop()) // Вывод: 2, true
    fmt.Println(stack.Pop()) // Вывод: 1, true
    fmt.Println(stack.Pop()) // Вывод: 0, false
}

