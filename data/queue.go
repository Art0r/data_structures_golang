package data

import "fmt"

type Queue struct {
	Head *node
	Rear *node
}

func (queue *Queue) Push(value interface{}) {
	if queue.Rear == nil {
		queue.Rear = &node{Data: value, Next: nil, Previous: nil}
		return
	}

	if queue.Head == nil {
		queue.Head = &node{Data: value, Next: queue.Rear, Previous: nil}
		queue.Rear.Previous = queue.Head
		return
	}

	node := &node{Data: value, Next: queue.Head, Previous: nil}

	node.Next.Previous = node
	node.Next.Previous.Next = queue.Head
	queue.Head = node

}

func (queue *Queue) Display() {

	s := queue.Head

	for s != nil {

		var next interface{}
		var previous interface{}

		if s.Next == nil {
			next = nil
		} else {
			next = s.Next.Data
		}

		if s.Previous == nil {
			previous = nil
		} else {
			previous = s.Previous.Data
		}

		fmt.Printf("NEXT: %d | CURRENT: %d | PREVIOUS: %d\n", next, s.Data, previous)
		fmt.Println("---------------------------------------------")

		s = s.Next
	}
}

func (queue *Queue) DisplayPrevious() {

	s := queue.Rear

	for s != nil {

		var next interface{}
		var previous interface{}

		if s.Next == nil {
			next = nil
		} else {
			next = s.Next.Data
		}

		if s.Previous == nil {
			previous = nil
		} else {
			previous = s.Previous.Data
		}

		fmt.Printf("NEXT: %d | CURRENT: %d | PREVIOUS: %d\n", next, s.Data, previous)
		fmt.Println("---------------------------------------------")

		s = s.Previous
	}

}
