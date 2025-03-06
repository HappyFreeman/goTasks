package main

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// binarySearch функция для выполнения бинарного поиска в отсортированном массиве.
// Возвращает индекс целевого значения, если оно присутствует в массиве, иначе возвращает -1.
func binarySearch[T Ordered](arr []T, target T) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := low + (high-low) / 2  // Используем (high-low)/2 для избежания переполнения

        // Проверяем, равен ли средний элемент целевому значению
        if arr[mid] == target {
            return mid
        }

        // Если целевое значение больше, игнорируем левую половину
        if arr[mid] < target {
            low = mid + 1
        } else {
            // Если целевое значение меньше, игнорируем правую половину
            high = mid - 1
        }
    }

    // Если элемент не найден, возвращаем -1
    return -1
}

func main() {
    arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
    target := 23
    result := binarySearch(arr, target)
    if result != -1 {
        fmt.Printf("Элемент найден на позиции: %d\n", result)
    } else {
        fmt.Println("Элемент не найден в массиве.")
    }
}
