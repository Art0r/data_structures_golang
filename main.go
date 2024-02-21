package main

import (
	dt "github.com/Art0r/data_structures/data"
)

func main() {

	// fmt.Println("-------------------------------------")
	// fmt.Println("STACK")
	// fmt.Println("-------------------------------------")

	// stack := dt.Stack{}

	// stack.Push(1)
	// stack.Push(2)
	// stack.Push(3)

	// stack.Display()

	// fmt.Println("-------------------------------------")

	// stack.Pop()

	// stack.Display()

	// fmt.Println("-------------------------------------")
	// fmt.Println("QUEUE")
	// fmt.Println("-------------------------------------")

	// queue := dt.Queue{}

	// fmt.Println(queue.Head)
	// fmt.Println(queue.Rear)

	// queue.Push(1)
	// queue.Push(2)
	// queue.Push(3)
	// queue.Push(4)

	// fmt.Println("-------------------------------------")

	// fmt.Println(queue.Head)
	// fmt.Println(queue.Rear)

	// queue.Display()

	// queue.DisplayPrevious()

	// Create a new graph
	graph := &dt.Graph{}

	// Add some nodes
	node1 := graph.AddNode(1)
	node2 := graph.AddNode(2)
	node3 := graph.AddNode(3)
	node4 := graph.AddNode(4)

	// Add edges between nodes
	graph.AddEdge(node1, node2)
	graph.AddEdge(node1, node3)
	graph.AddEdge(node2, node4)
	graph.AddEdge(node3, node4)

	// Print the graph
	graph.PrintGraph()
}
