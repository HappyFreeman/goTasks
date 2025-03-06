package main

import (
	"fmt"
)

func bubbleSort(arr []int) []int {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
    return arr
}

func main() {
    sample := []int{64, 34, 25, 12, 22, 11, 90}
    sorted := bubbleSort(sample)
    fmt.Println("Отсортированный массив:", sorted)
}
