package main

import (
	"fmt"
)

type Graph map[int][]int

// dfs реализует алгоритм поиска в глубину
func dfs(graph Graph, start int) {
    visited := make(map[int]bool)
    var recur func(v int)
    recur = func(v int) {
        if visited[v] {
            return
        }
        visited[v] = true
        fmt.Println(v)
        for _, n := range graph[v] {
            recur(n)
        }
    }
    recur(start)
}

// bfs реализует алгоритм поиска в ширину
func bfs(graph Graph, start int) {
    queue := []int{start}
    visited := make(map[int]bool)
    visited[start] = true

    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]
        fmt.Println(vertex)
        for _, n := range graph[vertex] {
            if !visited[n] {
                visited[n] = true
                queue = append(queue, n)
            }
        }
    }
}

func main() {
    graph := Graph{
        0: {1, 2},
        1: {2},
        2: {0, 3},
        3: {3},
    }

    fmt.Println("Starting DFS from vertex 2:")
    dfs(graph, 2)

    fmt.Println("\nStarting BFS from vertex 2:")
    bfs(graph, 2)
}
