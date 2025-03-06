package main

import "fmt"

type Node struct {
    Value int
    Next  *Node
}

type LinkedList struct {
    Head *Node
}

// Добавление элемента в начало списка
func (list *LinkedList) Prepend(value int) {
    newNode := &Node{Value: value, Next: list.Head}
    list.Head = newNode
}

// Печать всех элементов списка
func (list *LinkedList) Print() {
    for node := list.Head; node != nil; node = node.Next {
        fmt.Printf("%d -> ", node.Value)
    }
    fmt.Println("nil")
}

// Поиск значения в списке
func (list *LinkedList) Find(value int) *Node {
	for node := list.Head; node != nil; node = node.Next {
        if node.Value == value {
            return node
        }
    }
    return nil
}

// Удаление элемента из списка
func (list *LinkedList) Remove(value int) {
	if list.Head == nil {
        return
    }
    if list.Head.Value == value {
        list.Head = list.Head.Next
        return
    }
    for node := list.Head; node.Next != nil; node = node.Next {
        if node.Next.Value == value {
            node.Next = node.Next.Next // 3 -> 2 -> 1 | 3 ->-> 1
            return 
        }
    }
}

// Реверсирование списка
func (list *LinkedList) Reverse() {
	if list.Head == nil || list.Head.Next == nil {
        return
    }
    var prev, curr, next *Node
    curr = list.Head
    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    list.Head = prev
}

func main() {
    list := &LinkedList{}
    list.Prepend(1)
    list.Prepend(2)
    list.Prepend(3)
	list.Prepend(4)
    list.Print()  // Вывод: 3 -> 2 -> 1 -> nil
}
