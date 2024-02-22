package programs

import (
	"fmt"
	"strings"

	dt "github.com/Art0r/data_structures/data"
)

type Vendor struct {
	Name       string
	Ocuppation string
}

func setupGraphs() *dt.Graph {
	// Create a new graph
	graph := &dt.Graph{}

	// Add some nodes
	node1 := graph.AddNode(&Vendor{Name: "Maria", Ocuppation: "Vendedora de manga"})
	node2 := graph.AddNode(&Vendor{Name: "João", Ocuppation: "Vendedor de mamão"})
	node3 := graph.AddNode(&Vendor{Name: "José", Ocuppation: "Vendedor de abacate"})
	node4 := graph.AddNode(&Vendor{Name: "Helena", Ocuppation: "Vendedora de abacaxi"})
	graph.AddNode(&Vendor{Name: "Alice", Ocuppation: "Vendedora de abacaxi"})
	node6 := graph.AddNode(&Vendor{Name: "Bob", Ocuppation: "Vendedor de abacate"})
	node7 := graph.AddNode(&Vendor{Name: "Carmem", Ocuppation: "Vendedora de mamão"})

	// Add edges between nodes
	graph.AddEdge(node1, node2)
	graph.AddEdge(node1, node3)
	graph.AddEdge(node2, node4)
	graph.AddEdge(node3, node4)
	graph.AddEdge(node7, node2)
	graph.AddEdge(node6, node7)
	graph.AddEdge(node3, node1)

	return graph
}

func setupQueue(graph *dt.Graph) *dt.Queue {
	queue := dt.Queue{}

	queue.Push(graph.Nodes[2])
	queue.Push(graph.Nodes[4])
	queue.Push(graph.Nodes[5])

	return &queue
}

func executeQueue(queue *dt.Queue) []string {
	var path []string

	current := queue.Rear()

	for current != nil {

		node := current.Data.(*dt.GraphNode)
		vendor := node.Data.(*Vendor)

		if contains(path, vendor.Name) {
			current = current.Next
			continue
		}

		path = append(path, vendor.Name)

		if strings.Contains(vendor.Ocuppation, "manga") {
			break
		}

		for i := 0; i < len(node.Neighbors); i++ {
			queue.Push(node.Neighbors[i])
		}

		queue.Pop()

		current = current.Next
	}

	return path
}


func contains(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func RunVendor() {
	graph := setupGraphs()
	queue := setupQueue(graph)
	fmt.Println(executeQueue(queue))
}
