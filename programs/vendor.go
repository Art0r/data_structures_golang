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
	node1 := graph.AddNode(&Vendor{Name: "Alice", Ocuppation: "Vendedora de abacaxi"})
	node2 := graph.AddNode(&Vendor{Name: "Bob", Ocuppation: "Vendedor de mamão"})
	node3 := graph.AddNode(&Vendor{Name: "Peggy", Ocuppation: "Vendedor de abacate"})
	node4 := graph.AddNode(&Vendor{Name: "Claire", Ocuppation: "Vendedora de abacaxi"})
	node5 := graph.AddNode(&Vendor{Name: "Anuj", Ocuppation: "Vendedora de mamão"})
	node6 := graph.AddNode(&Vendor{Name: "Thom", Ocuppation: "Vendedora de manga"})
	node7 := graph.AddNode(&Vendor{Name: "Jonny", Ocuppation: "Vendedora de mamão"})

	// Add edges between nodes
	graph.AddEdge(node2, node5)
	graph.AddEdge(node2, node3)
	graph.AddEdge(node1, node3)
	graph.AddEdge(node4, node6)
	graph.AddEdge(node4, node7)

	return graph
}

func setupQueue(graph *dt.Graph) *dt.Queue {
	queue := dt.Queue{}

	queue.Push(graph.Nodes[0])
	queue.Push(graph.Nodes[1])
	queue.Push(graph.Nodes[3])

	return &queue
}

func executeQueue(queue *dt.Queue) []*Vendor {
	var path []*Vendor

	current := queue.Rear()

	for current != nil {

		node := current.Data.(*dt.GraphNode)
		vendor := node.Data.(*Vendor)

		if contains(path, vendor) {
			current = current.Next
			continue
		}

		path = append(path, vendor)

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

func contains(list []*Vendor, item *Vendor) bool {
	for _, i := range list {
		if i.Name == item.Name && i.Ocuppation == item.Ocuppation {
			return true
		}
	}
	return false
}

func RunVendor() {
	graph := setupGraphs()
	queue := setupQueue(graph)

	path := executeQueue(queue)

	for i := 0; i < len(path); i++ {
		fmt.Println(path[i].Name)
		fmt.Println(path[i].Ocuppation)
		fmt.Println("--------------------------")
	}
}
