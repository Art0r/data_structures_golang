package programs

import (
	"fmt"
	"math"
)

type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

func (g *Graph) AddVertex(data interface{}) *Vertex {
	vertex := &Vertex{Data: data}
	g.Vertices = append(g.Vertices, vertex)
	return vertex
}

func (g *Graph) AddEdge(from, to *Vertex, weight int) {
	edge := &Edge{From: from, To: to, Weight: weight}
	g.Edges = append(g.Edges, edge)
}

type DijkstraData struct {
	Distances map[*Vertex]int
	Visited   map[*Vertex]bool
}

func (dd *DijkstraData) Init() {
	dd.Distances = make(map[*Vertex]int)
	dd.Visited = make(map[*Vertex]bool)
}

func (g *Graph) dijkstra(start *Vertex) map[*Vertex]int {

	data := DijkstraData{}
	data.Init()

	for _, vertex := range g.Vertices {
		if vertex != start {
			data.Distances[vertex] = math.MaxInt32
		}
	}
	data.Distances[start] = 0

	for {
		vertex := g.minDistance(data)
		if vertex == nil {
			break
		}
		data.Visited[vertex] = true

		for _, edge := range g.Edges {
			if edge.From == vertex && !data.Visited[edge.To] {
				newDist := data.Distances[vertex] + edge.Weight
				if newDist < data.Distances[edge.To] {
					data.Distances[edge.To] = newDist
				}
			}
		}
	}

	return data.Distances
}

func (g *Graph) minDistance(args DijkstraData) *Vertex {
	min := math.MaxInt32
	var minVertex *Vertex
	for _, vertex := range g.Vertices {
		if !args.Visited[vertex] && args.Distances[vertex] < min {
			min = args.Distances[vertex]
			minVertex = vertex
		}
	}
	args.Visited[minVertex] = true
	return minVertex
}

func ExecDijkstra(start interface{}) {
	g := Graph{}

	v1 := g.AddVertex(0)
	v2 := g.AddVertex(1)
	v3 := g.AddVertex(2)
	v4 := g.AddVertex(3)
	v5 := g.AddVertex(4)

	g.AddEdge(v1, v2, 7)
	g.AddEdge(v2, v1, 7)

	g.AddEdge(v2, v3, 1)
	g.AddEdge(v3, v2, 1)

	g.AddEdge(v4, v1, 9)
	g.AddEdge(v1, v4, 9)

	g.AddEdge(v5, v4, 6)
	g.AddEdge(v4, v5, 6)

	result := g.dijkstra(v5)

	fmt.Println(result)
}

type Vertex struct {
	Data interface{}
}

type Edge struct {
	From   *Vertex
	To     *Vertex
	Weight int
}
