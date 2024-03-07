package programs

import (
	"fmt"
	"math"
)

type dijkstraData struct {
	Weights map[*vertex]int
	Visited map[*vertex]bool
}

func (dd *dijkstraData) Init() {
	dd.Weights = make(map[*vertex]int)
	dd.Visited = make(map[*vertex]bool)
}

type vertex struct {
	Data interface{}
}

type edge struct {
	From   *vertex
	To     *vertex
	Weight int
}

type graph struct {
	Vertices []*vertex
	Edges    []*edge
}

func (g *graph) AddVertex(data interface{}) *vertex {
	vertex := &vertex{Data: data}
	g.Vertices = append(g.Vertices, vertex)
	return vertex
}

func (g *graph) AddEdge(from, to *vertex, weight int) {
	edge := &edge{From: from, To: to, Weight: weight}
	g.Edges = append(g.Edges, edge)
}

func (g *graph) dijkstra(start *vertex) map[*vertex]int {

	data := dijkstraData{}
	data.Init()

	/* for every vertex except for the **starting vertex**
	weight must be infinite in the beggining because
	we still dont know if this other vertex is
	"recheable". Basically in this step we're popullating
	our **weights** variable, being it ordered or not
	*/
	for _, vertex := range g.Vertices {
		if vertex != start {
			data.Weights[vertex] = math.MaxInt32
		}
	}
	// the start always weight of 0
	data.Weights[start] = 0

	g.minorDistanceToNodes(&data)

	return data.Weights
}

func (g *graph) minorDistanceToNodes(data *dijkstraData) func() {
	vertex := g.minDistance(data)
	if vertex == nil {
		return func() {}
	}

	data.Visited[vertex] = true

	for _, edge := range g.Edges {
		if edge.From == vertex && !data.Visited[edge.To] {
			newDist := data.Weights[vertex] + edge.Weight
			if newDist < data.Weights[edge.To] {
				data.Weights[edge.To] = newDist
			}
		}
	}

	return g.minorDistanceToNodes(data)
}

func (g *graph) minDistance(args *dijkstraData) *vertex {
	min := math.MaxInt32
	var minVertex *vertex
	for _, vertex := range g.Vertices {
		if !args.Visited[vertex] && args.Weights[vertex] < min {
			min = args.Weights[vertex]
			minVertex = vertex
		}
	}
	args.Visited[minVertex] = true
	return minVertex
}

func ExecDijkstra() {
	g := graph{}

	a := g.AddVertex("A")
	b := g.AddVertex("B")
	c := g.AddVertex("C")
	d := g.AddVertex("D")
	e := g.AddVertex("E")

	g.AddEdge(a, b, 4)
	g.AddEdge(a, c, 2)
	g.AddEdge(b, d, 2)
	g.AddEdge(c, e, 5)
	g.AddEdge(c, b, 1)
	g.AddEdge(b, c, 3)
	g.AddEdge(e, d, 1)
	g.AddEdge(b, e, 3)
	g.AddEdge(c, d, 4)

	result := g.dijkstra(a)

	fmt.Println("-------------------")
	for index, vertex := range result {
		fmt.Println("VERTEX: ", index.Data)
		fmt.Println("DISTANCE:", vertex)
		fmt.Println("-------------------")
	}
}
