package main

import (
	"fmt"
	"sort"
)

type Edge struct {
    start, end, weight int
}

type DisjointSet struct {
    parent, rank []int
}

func NewDisjointSet(vertexCount int) *DisjointSet {
    parent := make([]int, vertexCount)
    rank := make([]int, vertexCount)
    for i := range parent {
        parent[i] = i
    }
    return &DisjointSet{parent, rank}
}

func (set *DisjointSet) Find(v int) int {
    if set.parent[v] != v {
        set.parent[v] = set.Find(set.parent[v])
    }
    return set.parent[v]
}

func (set *DisjointSet) Union(x, y int) {
    rootX := set.Find(x)
    rootY := set.Find(y)
    if rootX != rootY {
        if set.rank[rootX] > set.rank[rootY] {
            set.parent[rootY] = rootX
        } else if set.rank[rootX] < set.rank[rootY] {
            set.parent[rootX] = rootY
        } else {
            set.parent[rootY] = rootX
            set.rank[rootX]++
        }
    }
}

func KruskalMST(vertices int, edges []Edge) []Edge {
    sort.Slice(edges, func(i, j int) bool {
        return edges[i].weight < edges[j].weight
    })

    mst := make([]Edge, 0)
    set := NewDisjointSet(vertices)

    for _, e := range edges {
        if set.Find(e.start) != set.Find(e.end) {
            set.Union(e.start, e.end)
            mst = append(mst, e)
        }
    }

    return mst
}

func GetVertexCount(edges []Edge) int {
    maxVertex := 0
    for _, e := range edges {
        if e.start > maxVertex {
            maxVertex = e.start
        }
        if e.end > maxVertex {
            maxVertex = e.end
        }
    }
    return maxVertex + 1
}

func main() {
    edges := []Edge{
        {0, 1, 4}, {0, 2, 4}, {1, 2, 2}, {1, 3, 5}, {2, 3, 8}, {2, 4, 10}, {3, 4, 9}, {3, 5, 4}, {4, 5, 6},
    }
    mst := KruskalMST(GetVertexCount(edges), edges)
    fmt.Println("Edges in MST:")
    for _, e := range mst {
        fmt.Printf("%d - %d: %d\n", e.start, e.end, e.weight)
    }
}
