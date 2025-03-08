package main

import (
	"fmt"
	"math"
)

type Edge struct {
	to   int
	cost int
}

type Graph struct {
	vertices int
	edges    [][]Edge
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make([][]Edge, vertices),
	}
}

func (g *Graph) AddEdge(from, to, cost int) {
	g.edges[from] = append(g.edges[from], Edge{to, cost})
}

func (g *Graph) Dijkstra(source int) ([]int, [][]int) {
	minDist := make([]int, g.vertices)
	prev := make([]int, g.vertices)
	visited := make([]bool, g.vertices)

	for i := 0; i < g.vertices; i++ {
		minDist[i] = math.MaxInt32
		prev[i] = -1
	}
	minDist[source] = 0

	for i := 0; i < g.vertices; i++ {
		u := -1
		for v := 0; v < g.vertices; v++ {
			if !visited[v] && (u == -1 || minDist[v] < minDist[u]) {
				u = v
			}
		}
		if u == -1 {
			break
		}
		visited[u] = true

		for _, e := range g.edges[u] {
			if minDist[u]+e.cost < minDist[e.to] {
				minDist[e.to] = minDist[u] + e.cost
				prev[e.to] = u
			}
		}
	}

	paths := make([][]int, g.vertices)
	for i := 0; i < g.vertices; i++ {
		paths[i] = reconstructPath(prev, source, i)
	}

	return minDist, paths
}

func reconstructPath(prev []int, source, target int) []int {
	if target == source {
		return []int{source}
	}
	path := []int{}
	for at := target; at != -1; at = prev[at] {
		path = append([]int{at}, path...)
	}
	if len(path) == 0 || path[0] != source {
		return nil
	}
	return path
}

func main() {
	graph := NewGraph(5)
	graph.AddEdge(0, 1, 10)
	graph.AddEdge(0, 2, 3)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 2)
	graph.AddEdge(2, 1, 4)
	graph.AddEdge(2, 3, 8)
	graph.AddEdge(2, 4, 2)
	graph.AddEdge(3, 4, 7)
	graph.AddEdge(4, 3, 9)

	distances, paths := graph.Dijkstra(0)
	fmt.Println("Кратчайшие расстояния от вершины 0:")
	for i, d := range distances {
		fmt.Printf("До вершины %d: %d\n", i, d)
	}

	fmt.Println("\nМаршруты:")
	for i, p := range paths {
		if p == nil {
			fmt.Printf("Вершина %d недостижима от 0\n", i)
		} else {
			fmt.Printf("Путь до %d: %v\n", i, p)
		}
	}
}
