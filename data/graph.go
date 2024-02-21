package data

import "fmt"

type GraphNode struct {
	Data      interface{}
	Neighbors []*GraphNode
}

type Graph struct {
	Nodes []*GraphNode
}

func (graph *Graph) AddNode(data interface{}) *GraphNode {
	node := &GraphNode{Data: data}
	graph.Nodes = append(graph.Nodes, node)
	return node
}

func (graph *Graph) AddEdge(node1, node2 *GraphNode) {
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)
}

func (graph *Graph) PrintGraph() {
	for _, node := range graph.Nodes {
		fmt.Printf("Node %d neighbors: ", node.Data)
		for _, neighbor := range node.Neighbors {
			fmt.Printf("%d ", neighbor.Data)
		}
		fmt.Println()
	}
}
