package main

import "fmt"

// Graph структура, представляющая граф
type Graph struct {
    vertices []*Vertex
}

// Vertex структура, представляющая вершину графа
type Vertex struct {
    Key int
    Adjacent []*Vertex
}

// AddVertex добавляет вершину в граф
func (g *Graph) AddVertex(k int) {
    if contains(g.vertices, k) {
        err := fmt.Errorf("vertex %v not added because it is an existing key", k)
        fmt.Println(err.Error())
        return
    }
    g.vertices = append(g.vertices, &Vertex{Key: k})
}

// AddEdge добавляет ребро между двумя вершинами в графе
func (g *Graph) AddEdge(from, to int) {
    // Получаем вершины
    fromVertex := g.getVertex(from)
    toVertex := g.getVertex(to)
    // Проверяем, что вершины существуют
    if fromVertex == nil || toVertex == nil {
        err := fmt.Errorf("invalid edge (%v->%v)", from, to)
        fmt.Println(err.Error())
        return
    }
    // Проверяем, что ребро не существует
    if contains(fromVertex.Adjacent, to) {
        err := fmt.Errorf("existing edge (%v->%v)", from, to)
        fmt.Println(err.Error())
        return
    }
    // Добавляем ребро
    fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
}

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
    g.dfsRecursive(start, visited)
	fmt.Println()
}


func (g *Graph) dfsRecursive(v int, visited map[int]bool) {
	vertex := g.getVertex(v)
	if vertex == nil || visited[v] {
		return
	}
	visited[v] = true
	fmt.Print(vertex.Key, " ")
	for _, neighbor := range vertex.Adjacent {
        g.dfsRecursive(neighbor.Key, visited)
    }
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
    queue := []int{start}
    for len(queue) > 0 {
        vertex := g.getVertex(queue[0])
        queue = queue[1:]
        if vertex == nil || visited[vertex.Key] {
            continue
        }
        visited[vertex.Key] = true
        fmt.Print(vertex.Key, " ")
        for _, neighbor := range vertex.Adjacent {
            queue = append(queue, neighbor.Key)
        }
    }
    fmt.Println()
}

func (g *Graph) Print() {
	 // Выводим граф
	 for _, v := range g.vertices {
        fmt.Printf("Vertex %v : ", v.Key)
        for _, v := range v.Adjacent {
            fmt.Printf("%v ", v.Key)
        }
        fmt.Println()
    }
}

// Помощник для проверки существования вершины
func contains(s []*Vertex, k int) bool {
    for _, v := range s {
        if k == v.Key {
            return true
        }
    }
    return false
}

// Помощник для получения вершины по ключу
func (g *Graph) getVertex(k int) *Vertex {
    for _, v := range g.vertices {
        if v.Key == k {
            return v
        }
    }
    return nil
}

func main() {
    graph := &Graph{}
    for i := 1; i <= 5; i++ {
        graph.AddVertex(i)
    }
    graph.AddEdge(1, 2)
    graph.AddEdge(1, 3)
    graph.AddEdge(2, 3)
    graph.AddEdge(4, 2)
    graph.AddEdge(5, 2)
    graph.AddEdge(3, 5)

	graph.Print()
	graph.DFS(3)
}
